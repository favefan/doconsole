package v1

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/service"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	}

	result, err := service.ContainerCreate(body)
	if err != nil {
		log.Println(err)
	}

	appG.Response(http.StatusOK, e.Success, result)
}
