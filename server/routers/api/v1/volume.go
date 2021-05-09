package v1

import (
	"context"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VolumeList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
		return
	}
	defer cli.Close()

	filter := filters.Args{}
	volumes, err := cli.VolumeList(ctx, filter)
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, volumes)
	return
}

func VolumeCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	name := json["Name"].(string)
	driver := json["Driver"].(string)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
		return
	}
	defer cli.Close()

	
	volumes, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Driver:     driver,
		DriverOpts: nil,
		Labels:     nil,
		Name:       name,
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

	appG.Response(http.StatusOK, e.Success, volumes)
	return
}

func VolumeInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	name := c.Param("name")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
		return
	}
	defer cli.Close()

	volume, err := cli.VolumeInspect(ctx, name)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, volume)
	return
}

func VolumesRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	var arr []string
	removeListOfFail := arr[:]

	json := make(map[string][]string) //注意该结构接受的内容
	c.BindJSON(&json)
	// fmt.Println(json)

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		appG.Response(
			http.StatusInternalServerError,
			e.ErrorDockerDaemonConnectionFailed,
			nil,
		)
		return
	}

	defer cli.Close()


	for _, v := range  json["array"] {
		if err := cli.VolumeRemove(ctx, v, true); err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}

	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}


