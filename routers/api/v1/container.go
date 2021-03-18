package v1

import (
	"log"
	"net/http"
	_ "strconv"

	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/pkg/util"
	"gitee.com/favefan/doconsole/services/containerservice"
	"github.com/docker/docker/api/types/container"
	_ "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// ListContainers returns the list of containers in the docker host.
func ListContainers(c *gin.Context) {
	appG := app.Gin{C: c}

	containers, err := containerservice.List()
	switch {
	case client.IsErrConnectionFailed(err):
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
		return
	case err == nil:
	default:
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, containers)
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

func StartContainer() {}

func StopContainer() {}

func RestartContainer() {}

func KillContainer() {}

func PauseContainer() {}

func ResumeContainer() {}

func RemoveContainer() {}

func DeleteStoppedContainers() {}
