package db

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"facemasq/lib/logging"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type MySQLContainer struct {
	ID         string
	Connection ConnectionParams
	Cleanup    func() error
	dbType     string
}

func StartMySQLContainer(dbName, dbUser, dbPassword, dbPort string) (testContainer *MySQLContainer, err error) {
	var cntnr container.ContainerCreateCreatedBody
	var containerList []types.Container
	var logsReader io.ReadCloser
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	found := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == "mysql:8-debian" {
				found = true
			}
		}
	}

	if !found {
		_, err := cli.ImagePull(
			context.Background(),
			"docker.io/library/mysql:8-debian",
			types.ImagePullOptions{},
		)
		if err != nil {
			panic(err)
		}
	}

	filters := filters.NewArgs()
	filters.Add("name", "mysql_facemasq")

	containerList, err = cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: filters})
	if err != nil {
		return
	}

	if len(containerList) > 0 {
		for i := range containerList {
			fmt.Printf("%+v", containerList[i])
		}
	}
	// Prepare container
	cntnr, err = cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "mysql:8-debian",
			Cmd:   []string{},
			Env: []string{
				fmt.Sprintf("MYSQL_DATABASE=%s", dbName),
				fmt.Sprintf("MYSQL_USER=%s", dbUser),
				fmt.Sprintf("MYSQL_PASSWORD=%s", dbPassword),
				fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", dbPassword),
			},
		},
		&container.HostConfig{
			PublishAllPorts: true,
			AutoRemove:      true,
		},
		&network.NetworkingConfig{},
		&v1.Platform{
			Architecture: os.Getenv("GOARCH"),
			OS:           "linux",
		},
		"mysql_facemasq",
	)
	if err != nil {
		cli.Close()
		return
	}

	cleanup := func() error {
		if err := cli.ContainerKill(context.Background(), cntnr.ID, ""); err != nil {
			return err
		}
		// if err := cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{RemoveVolumes: true, RemoveLinks: true, Force: true}); err != nil {
		// 	return err
		// }
		return cli.Close()
	}

	// Start container
	err = cli.ContainerStart(context.Background(), cntnr.ID, types.ContainerStartOptions{})
	if err != nil {
		_ = cleanup()
		return
	}

	// Inspect the container so we can get published port for container
	inspect, err := cli.ContainerInspect(context.Background(), cntnr.ID)
	if err != nil {
		_ = cleanup()
		return
	}

	// Get the container's network ports from the inspector
	bindings, ok := inspect.NetworkSettings.Ports[nat.Port(dbPort+"/tcp")]
	if !ok || len(bindings) < 1 {
		_ = cleanup()
		return
	}

	// Read the container's logs and scan them for a message that indicates the container has sucessessfully started and is ready for connections
	logsReader, err = cli.ContainerLogs(context.Background(), cntnr.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		_ = cleanup()
		return nil, err
	}
	logsScanner := bufio.NewScanner(logsReader)
	logsScanner.Split(bufio.ScanLines)
	done := make(chan error)

	go func() {
		for logsScanner.Scan() {
			text := logsScanner.Text()
			if strings.Contains(text, "mysqld") && strings.Contains(text, fmt.Sprintf("port: %s", dbPort)) {
				done <- nil
				break
			}
		}
	}()
	select {
	case err := <-done:
		if err != nil {
			_ = cleanup()
			return nil, err
		}
	case <-time.After(time.Second * 100):
		_ = cleanup()
		return nil, fmt.Errorf("timeout waiting for container")
	}
	logging.Debug1("Container is running")
	testContainer = &MySQLContainer{
		ID: cntnr.ID,
		Connection: ConnectionParams{
			DBName: dbName,
			DBUser: dbUser,
			DBPass: dbPassword,
			DBHost: "localhost",
			DBPort: bindings[0].HostPort,
		},
		Cleanup: cleanup,
		dbType:  "mysql",
	}
	return
}

func (mysql *MySQLContainer) Close() error {
	return mysql.Cleanup()
}

func (mysql *MySQLContainer) GetConnection() ConnectionParams {
	return mysql.Connection
}
