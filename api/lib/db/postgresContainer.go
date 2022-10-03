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

type Postgres struct {
	ID         string
	Connection ConnectionParams
	Cleanup    func() error
}

func StartPostgresDB(dbName, dbUser, dbPassword, dbPort string) (testContainer *Postgres, err error) {
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

	// Create container
	container, err := cli.ContainerCreate(
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
		if err := cli.ContainerKill(context.Background(), container.ID, ""); err != nil {
			return err
		}
		return cli.Close()
	}

	// Start container
	err = cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		_ = cleanup()
		return
	}

	// Get published port for postgres
	inspect, err := cli.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		_ = cleanup()
		return
	}
	// bindings[0].HostIP + ":" + bindings[0].HostPort
	bindings, ok := inspect.NetworkSettings.Ports[nat.Port(dbPort+"/tcp")]
	if !ok || len(bindings) < 1 {
		_ = cleanup()
		return
	}
	// Monitor Logs for readiness
	logsReader, err := cli.ContainerLogs(context.Background(), container.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true})
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
			// Posstgres DB starts twice to get really ready
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

	testContainer = &Postgres{
		ID: container.ID,
		Connection: ConnectionParams{
			DBName: dbName,
			DBUser: dbUser,
			DBPass: dbPassword,
			DBHost: "localhost",
			DBPort: bindings[0].HostPort,
		},
		Cleanup: cleanup,
	}
	return
}

func (postgres *Postgres) Close() error {
	return postgres.Cleanup()
}

func (postgres *Postgres) GetConnection() ConnectionParams {
	return postgres.Connection
}
