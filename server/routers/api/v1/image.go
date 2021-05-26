package v1

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ImageList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	images, err := global.GClient.ImageList(ctx, types.ImageListOptions{
		All:     true,
		Filters: filters.Args{},
	})
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, images)
	return
}

func ImageInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	imageID := c.Param("id")

	image, _, err := global.GClient.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, image)
	return
}

func ImageHistory(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	imageID := c.Param("id")

	imageHistory, err := global.GClient.ImageHistory(ctx, imageID)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, imageHistory)
	return
}

func ImagesRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	var arr []string
	removeListOfFail := arr[:]

	json := make(map[string][]string) //注意该结构接受的内容
	c.BindJSON(&json)

	for _, v := range  json["array"] {
		_, err := global.GClient.ImageRemove(ctx, v, types.ImageRemoveOptions{
			Force:         false,
			PruneChildren: false,
		})
		if err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}
	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}

func ImageSearch(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	term := c.Query("term")
	//limit := c.Query("limit")

	searchResult, err := global.GClient.ImageSearch(ctx, term, types.ImageSearchOptions{
		RegistryAuth:  "",
		PrivilegeFunc: nil,
		Filters:       filters.Args{},
		Limit:         50,
	})
	if err != nil {
		fmt.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, searchResult)
	return
}

func ImageInspectFromRegistry(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	imageID := c.Param("id")
	//limit := c.Query("limit")

	inspectResult, err := global.GClient.DistributionInspect(ctx, imageID, "")
	if err != nil {
		fmt.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, inspectResult)
	return
}

func ImagePull(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	image := c.Query("ref")

	reader, err := global.GClient.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			"Image pulling failed.",
		)
		return
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	fmt.Println(buf)

	appG.Response(http.StatusOK, e.Success, "ok")
	return
}
