package db

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"facemasq/lib/logging"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type PostgreSQLContainer struct {
	ID         string
	Connection ConnectionParams
	Cleanup    func() error
	dbType     string
}

func StartPostgreSQLContainer(dbName, dbUser, dbPassword, dbPort string) (testContainer *PostgreSQLContainer, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return
	}

	found := false
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == "postgres:14-alpine" {
				found = true
			}
		}
	}

	if !found {
		if _, err = cli.ImagePull(
			context.Background(),
			"docker.io/library/postgres:14-alpine",
			types.ImagePullOptions{},
		); err != nil {
			return
		}
	}

	// Prepare container
	cntnr, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "postgres:14-alpine",
			Cmd:   []string{},
			Env: []string{
				fmt.Sprintf("POSTGRES_DB=%s", dbName),
				fmt.Sprintf("POSTGRES_USER=%s", dbUser),
				fmt.Sprintf("POSTGRES_PASSWORD=%s", dbPassword),
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
		"postgres_facemasq",
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
	logsReader, err := cli.ContainerLogs(context.Background(), cntnr.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
	if err != nil {
		_ = cleanup()
		return
	}
	logsScanner := bufio.NewScanner(logsReader)
	logsScanner.Split(bufio.ScanLines)
	done := make(chan error)

	go func() {
		matchCount := 0
		for logsScanner.Scan() {
			text := logsScanner.Text()
			// PostgreSQL DB starts twice to get really ready
			if strings.Contains(text, "database system is ready to accept connections") {
				matchCount++
				if matchCount == 2 {
					done <- nil
					break
				}
			}
		}
	}()
	select {
	case err = <-done:
		if err != nil {
			_ = cleanup()
			return
		}
	case <-time.After(time.Second * 100):
		_ = cleanup()
		err = fmt.Errorf("timeout waiting for postgres")
		return
	}
	logging.Debug("Container is running")
	testContainer = &PostgreSQLContainer{
		ID: cntnr.ID,
		Connection: ConnectionParams{
			DBName: dbName,
			DBUser: dbUser,
			DBPass: dbPassword,
			DBHost: "localhost",
			DBPort: bindings[0].HostPort,
		},
		Cleanup: cleanup,
		dbType:  "postgresql",
	}
	return
}

func (postgres *PostgreSQLContainer) Close() error {
	return postgres.Cleanup()
}

func (postgres *PostgreSQLContainer) GetConnection() ConnectionParams {
	return postgres.Connection
}
