package service

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/go-connections/nat"
	"log"
)

type ContainerCreationBody struct {
	Name string `form:"Name"`
	Image string `form:"Image"`
	ExposedPorts nat.PortSet `form:"ExposedPorts"`
	Env []string `form:"Env"`
	Cmd []string `form:"Cmd"`
	Entrypoint strslice.StrSlice `form:"Entrypoint"`
	WorkingDir string `form:"WorkingDir"`
	User string `form:"User"`
	Hostname string `form:"Hostname"`
	Domainname string `form:"Domainname"`
	MacAddress string `form:"MacAddress"`
	Tty bool `form:"Tty"`
	OpenStdin bool `form:"OpenStdin"`
	HostConfig struct{
		Binds []string `form:"HostConfig.Binds"`
		PortBindings nat.PortMap `form:"HostConfig.PortBindings"`
		PublishAllPorts bool `form:"HostConfig.PublishAllPorts"`
		AutoRemove bool `form:"HostConfig.AutoRemove"`
		RestartPolicy struct {
			Name string `form:"HostConfig.RestartPolicy.Name"`
		}
		NetworkMode container.NetworkMode `form:"HostConfig.NetworkMode"`
	}
	Volumes map[string]struct{} `form:"Volumes"`
	NetworkingConfig struct {
		EndpointsConfig map[string]*network.EndpointSettings `form:"NetworkingConfig.EndpointsConfig"`
	}
}

type ContainerLogsOptions struct {
	ShowStdout bool `form:"ShowStdout"`
	ShowStderr bool `form:"ShowStderr"`
	Since string `form:"Since"`
	Until string `form:"Until"`
	Timestamps bool `form:"Timestamps"`
	Follow bool `form:"Follow"`
	Tail string `form:"Tail"`
	Details bool `form:"Details"`
}

type ExecConfig struct {
	User         string `form:"User"`   // User that will run the command
	// Privileged   bool `form:"Privileged"`    // Is the container in privileged mode
	Tty          bool `form:"Tty"`    // Attach standard streams to a tty.
	AttachStdin  bool `form:"AttachStdin"`    // Attach the standard input, makes possible user interaction
	AttachStderr bool `form:"AttachStderr"`    // Attach the standard error
	AttachStdout bool `form:"AttachStdout"`    // Attach the standard output
	// Detach       bool `form:"Detach"`    // Execute in detach mode
	// DetachKeys   string `form:"DetachKeys"`   // Escape keys for detach
	// Env          []string `form:"Env"` // Environment variables
	// WorkingDir   string `form:"WorkingDir"`   // Working directory
	Cmd          []string `form:"Cmd"` // Execution commands and args
}

type ExecAttachConfig struct {
	ExecId string `form:"ExecId"`
	Detach bool `form:"Detach"`
	Tty    bool `form:"Tty"`
}

func ContainerCreate(body ContainerCreationBody) (string, error) {
	ctx := context.Background()
	// fmt.Println(body.NetworkingConfig.EndpointsConfig["bridge"].IPAMConfig)

	result, err := global.GClient.ContainerCreate(ctx,
		&container.Config{
			Hostname:        body.Hostname,
			Domainname:      body.Domainname,
			User:            body.User,
			ExposedPorts:    body.ExposedPorts,
			Tty:             body.Tty,
			OpenStdin:       body.OpenStdin,
			Env:             body.Env,
			Cmd:             body.Cmd,
			Image:           body.Image,
			Volumes:         body.Volumes,
			WorkingDir:      body.WorkingDir,
			Entrypoint:      body.Entrypoint,
			MacAddress:      body.MacAddress,
		},
		&container.HostConfig{
			Binds: body.HostConfig.Binds,
			PortBindings: body.HostConfig.PortBindings,
			PublishAllPorts: body.HostConfig.PublishAllPorts,
			AutoRemove: body.HostConfig.AutoRemove,
			RestartPolicy: container.RestartPolicy{
				Name: body.HostConfig.RestartPolicy.Name,
			},
			NetworkMode: body.HostConfig.NetworkMode,
		},
		&network.NetworkingConfig{
			EndpointsConfig: body.NetworkingConfig.EndpointsConfig,
		}, nil, body.Name)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return result.ID, nil
}

func ContainerExecAttach(execID string) (hijackedResponse types.HijackedResponse, err error) {
	ctx := context.Background()

	hijackedResponse, err = global.GClient.ContainerExecAttach(ctx, execID, types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	return

}
