//go:build database

package db

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type MySQL struct {
	ID         string
	Connection ConnectionParams
	Cleanup    func() error
}

func StartMySQLDB(dbName, dbUser, dbPassword, dbPort string) (*MySQL, error) {
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
		if _, err := cli.ImagePull(
			context.Background(),
			"docker.io/library/mysql:8-debian",
			types.ImagePullOptions{},
		); err != nil {
			panic(err)
		}
	}

	// Create container
	container, err := cli.ContainerCreate(
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
		return nil, err
	}

	cleanup := func() error {
		if err := cli.ContainerKill(context.Background(), container.ID, ""); err != nil {
			return err
		}
		return cli.Close()
	}

	// Start container
	err = cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		_ = cleanup()
		return nil, err
	}

	// Get published port for container
	inspect, err := cli.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		_ = cleanup()
		return nil, err
	}
	// bindings[0].HostIP + ":" + bindings[0].HostPort

	bindings, ok := inspect.NetworkSettings.Ports[nat.Port(dbPort+"/tcp")]
	if !ok || len(bindings) < 1 {
		_ = cleanup()
		return nil, err
	}
	// Monitor Logs for readiness
	logsReader, err := cli.ContainerLogs(context.Background(), container.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
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
	fmt.Println("Container is running")
	return &MySQL{
		ID: container.ID,
		Connection: ConnectionParams{
			DBName: dbName,
			DBUser: dbUser,
			DBPass: dbPassword,
			DBHost: "localhost",
			DBPort: bindings[0].HostPort,
		},
		Cleanup: cleanup,
	}, nil
}

func (mysql *MySQL) Close() error {
	return mysql.Cleanup()
}

func (mysql *MySQL) GetConnection() ConnectionParams {
	return mysql.Connection
}
