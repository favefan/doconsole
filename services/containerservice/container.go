package containerservice

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Container struct {
	ContainerID   string
	ContainerName string

	Config           container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
	Platform         v1.Platform
}

func (c *Container) Create() (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()

	// Create API client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	defer cli.Close()

	// Image exits check

	createdBody, err := cli.ContainerCreate(
		ctx,
		&c.Config,
		&c.HostConfig,
		nil,
		nil,
		c.ContainerName,
	)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	return createdBody, nil
}

type ContainerOperationOptions struct {
	ListOptions   types.ContainerListOptions
	AttachOptions types.ContainerAttachOptions
	CommitOptions types.ContainerCommitOptions
	LogsOptions   types.ContainerLogsOptions
	RemoveOptions types.ContainerRemoveOptions
	StartOptions  types.ContainerStartOptions
}
