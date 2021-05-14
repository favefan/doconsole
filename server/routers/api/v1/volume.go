package v1

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func VolumeList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	filter := filters.Args{}
	volumes, err := global.GClient.VolumeList(ctx, filter)
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

func VolumeCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	name := json["Name"].(string)
	driver := json["Driver"].(string)

	volumes, err := global.GClient.VolumeCreate(ctx, volume.VolumeCreateBody{
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

	volume, err := global.GClient.VolumeInspect(ctx, name)
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

	for _, v := range  json["array"] {
		if err := global.GClient.VolumeRemove(ctx, v, true); err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}

	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}


