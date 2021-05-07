package v1

import (
	"context"
	"fmt"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func ImageList(c *gin.Context) {
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

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     true,
		Filters: filters.Args{},
	})
	if err != nil {
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

	image, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
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

	imageHistory, err := cli.ImageHistory(ctx, imageID)
	if err != nil {
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

	_, err = cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: false,
	})
	if err != nil {
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

	searchResult, err := cli.ImageSearch(ctx, term, types.ImageSearchOptions{
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

	inspectResult, err := cli.DistributionInspect(ctx, imageID, "")
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

	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{
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
