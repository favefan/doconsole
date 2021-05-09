package v1

import (
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DaemonHost(c *gin.Context) {
	appG := app.Gin{C: c}

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

	addr := cli.DaemonHost()
	appG.Response(
		http.StatusOK,
		e.Success,
		addr,
		)
	return
}
