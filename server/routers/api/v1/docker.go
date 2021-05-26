package v1

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/initialize"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func CreateClient(c *gin.Context) {
	appG := app.Gin{C: c}
	host := initialize.Host{}
	var cli *client.Client
	var err error

	if err = c.ShouldBindJSON(&host); err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			"参数错误")
		return
	}
	cli, err = initialize.CreateDockerClient(&host)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err)
		return
	}
	global.GClient = cli
	appG.Response(
		http.StatusOK,
		e.Success,
		host.Name)
	return
}

func DaemonHost(c *gin.Context) {
	appG := app.Gin{C: c}

	addr := global.GClient.DaemonHost()
	appG.Response(
		http.StatusOK,
		e.Success,
		addr,
		)
	return
}

func ClientVersion(c *gin.Context) {
	appG := app.Gin{C: c}
	version := global.GClient.ClientVersion()
	appG.Response(
		http.StatusOK,
		e.Success,
		version)
	return
}

func Info(c *gin.Context) {
	ctx := context.Background()
	appG := app.Gin{C: c}

	info, err := global.GClient.Info(ctx)
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			err,
		)
		return
	}
	appG.Response(
		http.StatusOK,
		e.Success,
		info,
	)
	return
}