package v1

import (
	"context"
	"gitee.com/favefan/doconsole/global"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NetworkList(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	networks, err := global.GClient.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(http.StatusOK, e.Success, networks)
	return
}

func NetworkInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	id := c.Param("id")

	networks, err := global.GClient.NetworkInspect(ctx, id, types.NetworkInspectOptions{
		Scope:   "",
		Verbose: true,
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

	appG.Response(http.StatusOK, e.Success, networks)
	return
}

//type networkCreateBody struct {
//	name string `form:"Name"`
//	driver string `form:"Driver"`
//	attachable bool `form:"Attachable"`
//	internal string `form:"Internal"`
//	labels struct{
//		label string `form:"Label"`
//	}
//}

func NetworkCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)

	attachable := json["Attachable"].(bool)
	internal := json["Internal"].(bool)
	// label := c.PostForm(" Label")

	// fmt.Println(containerService)
	result, err := global.GClient.NetworkCreate(ctx, json["Name"].(string), types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         json["Driver"].(string),
		Scope:          "local",
		EnableIPv6:     false,
		IPAM: 			&network.IPAM{
			Driver: "default",
			Options: nil,
			Config: nil,
		},
		Internal:       internal,
		Attachable:     attachable,
		Ingress:        false,
		ConfigOnly:     false,
		ConfigFrom:     nil,
		Options:        nil,
		Labels:         nil,
	})
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}
	appG.Response(http.StatusOK, e.Success, result)
	return
}

func NetworksRemove(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()
	var arr []string
	removeListOfFail := arr[:]

	json := make(map[string][]string) //注意该结构接受的内容
	c.BindJSON(&json)
	// fmt.Println(json)

	for _, v := range  json["array"] {
		if err := global.GClient.NetworkRemove(ctx, v); err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}

	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}

func NetworkConnect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	networkID := c.Param("id")
	json := make(map[string]string)
	c.BindJSON(&json)

	err := global.GClient.NetworkConnect(ctx, networkID, json["Container"], nil)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	appG.Response(http.StatusOK, e.Success, "ok")
	return
}

func NetworkDisconnect(c *gin.Context) {
	appG := app.Gin{C: c}
	ctx := context.Background()

	networkID := c.Param("id")
	json := make(map[string]string)
	c.BindJSON(&json)

	err := global.GClient.NetworkDisconnect(ctx, networkID, json["Container"], true)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, err)
		return
	}

	appG.Response(http.StatusOK, e.Success, "ok")
	return
}
