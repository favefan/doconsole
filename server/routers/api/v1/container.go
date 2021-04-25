package v1

import (
	"context"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/pkg/util"
	"gitee.com/favefan/doconsole/services/containerservice"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "strconv"
)

// ContainerList returns the list of containers in the docker host.
func ContainerList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx,
		types.ContainerListOptions{
			All:  true,
			Size: true,
		})
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, containers)
	return
}

// ContainerInspect
func ContainerInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Param("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	container, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, container)
	return
}

func ContainerStart(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerStop(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerPause(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	err = cli.ContainerPause(ctx, containerID)
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerRestart(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	err = cli.ContainerRestart(ctx, containerID, nil)
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
	}
	defer cli.Close()

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

type ContainerCreateBody struct {
	ContainerName string `form:"container_name"`
	AutoPull      bool   `form:"auto_pull"`

	Config struct {
		HostName   string            `form:"host_name"`
		DomainName string            `form:"domain_name"`
		Env        []string          `form:"env"`
		Cmd        strslice.StrSlice `form:"cmd"`
		Shell      strslice.StrSlice `form:"shell"`
		Image      string            `form:"image" binding:"required"`
	}

	HostConfig struct {
		NetworkMode   container.NetworkMode `form:"network_mode"`
		RestartPolicy struct {
			Name string `form:"restart"`
			// MaximumRetryCount int
		}
		AutoRemove bool `form:"auto_remove"`

		// Applicable to UNIX platforms
		PublishAllPorts bool `form:"publish_all"`
	}
}

// CreateContainer create a container with given parameters.
func CreateContainer(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		createBody ContainerCreateBody
	)

	if err := c.ShouldBind(&createBody); err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	containerService := new(containerservice.Container)
	if err := util.StructAdaptiveCopyByJSON(&createBody, containerService); err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(containerService)
	result, err := containerService.Create()
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, result)
	return
}

func InspectContainer() {}

func ListContainerProcesses() {}

func GetContainerLogs() {}

func StopContainer() {}

func RestartContainer() {}

func KillContainer() {}

func PauseContainer() {}

func ResumeContainer() {}

func RemoveContainer() {}

func DeleteStoppedContainers() {}
