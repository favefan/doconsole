package v1

import (
	"context"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"net/http"
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
	}

	appG.Response(http.StatusOK, e.Success, images)
	return
}
