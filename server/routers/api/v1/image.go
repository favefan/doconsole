package v1

import (
	"context"
	"fmt"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
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
			nil,
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
			nil,
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
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, imageHistory)
	return
}

func ImageRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	imageID := c.Param("id")

	_, err := global.GClient.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: false,
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
			nil,
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

	reader, err := global.GClient.ImagePull(ctx, image, types.ImagePullOptions{
		All:           false,
		RegistryAuth:  "",
		PrivilegeFunc: nil,
		Platform:      "",
	})
	if err != nil {
		fmt.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	appG.Response(http.StatusOK, e.Success, "inspectResult")
	return
}
