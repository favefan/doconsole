package v1

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/service"
	"gitee.com/favefan/doconsole/service/websocket"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	_ "strconv"
)

// ContainerList returns the list of containers in the docker host.
func ContainerList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	containers, err := global.GClient.ContainerList(ctx,
		types.ContainerListOptions{
			All:  true,
			Size: true,
		})
	if err != nil {
		log.Println(err)
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

func ContainerInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Param("id")

	container, err := global.GClient.ContainerInspect(ctx, containerID)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, container)
	return
}

func ContainerStart(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	err := global.GClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerStop(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	err := global.GClient.ContainerStop(ctx, containerID, nil)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerPause(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	err := global.GClient.ContainerPause(ctx, containerID)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerRestart(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Query("id")

	err := global.GClient.ContainerRestart(ctx, containerID, nil)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Param("id")

	err := global.GClient.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         false,
	})
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, nil)
	return
}

func ContainerCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	// ctx := context.Background()

	var body service.ContainerCreationBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, err)
		return
	}

	result, err := service.ContainerCreate(body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	appG.Response(http.StatusOK, e.Success, result)
	return
}

func ContainerUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Param("id")
	json := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&json)
	_, err := global.GClient.ContainerUpdate(ctx, containerID, container.UpdateConfig{
		RestartPolicy: container.RestartPolicy{
			Name: json["RestartPolicy"],
		},
	})
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	appG.Response(http.StatusOK, e.Success, "ok")
	return
}

func ContainerLogs(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	containerID := c.Param("id")
	var options service.ContainerLogsOptions
	err := c.ShouldBindJSON(&options)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, err)
		return
	}

	logs, err := global.GClient.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
		ShowStdout: options.ShowStdout,
		ShowStderr: options.ShowStderr,
		Since:      options.Since,
		Until:      options.Until,
		Timestamps: options.Timestamps,
		Follow:     options.Follow,
		Tail:       options.Tail,
		Details:    options.Details,
	})
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}
	defer logs.Close()
	io.Copy(os.Stdout, logs)

}

func ContainerExecCreate(c *gin.Context) {
	ctx := context.Background()
	appG := app.Gin{C: c}

	container := c.Param("id")
	var config service.ExecConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, err)
		return
	}

	result, err := global.GClient.ContainerExecCreate(ctx, container, types.ExecConfig{
		User:         config.User,
		Tty:          config.Tty,
		AttachStdin:  config.AttachStdin,
		AttachStderr: config.AttachStderr,
		AttachStdout: config.AttachStdout,
		Cmd:          config.Cmd,
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}
	appG.Response(http.StatusOK, e.Success, result)
}


func ContainerExecAttach(c *gin.Context) {
	container := c.Param("id")
	websocket.TerminalHandle(c, container)
}
