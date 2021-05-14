package v1

import (
	_ "context"
	"gitee.com/favefan/doconsole/models"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HostCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	//ctx := context.Background()

	bodyJson := make(map[string]interface{})
	c.BindJSON(&bodyJson)
	host := models.Host{
		Model: models.Model{},
		Name: bodyJson["Name"].(string),
		ViaSocket: bodyJson["ViaSocket"].(bool),
		DockerEngineURL: bodyJson["DockerEngineURL"].(string),
		HostIP: bodyJson["HostIP"].(string),
		TLS: bodyJson["TLS"].(bool),
	}

	if err := host.Create(); err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
			)
		return
	}

	appG.Response(
		http.StatusOK,
		e.Success,
		nil,
	)
	return

}

func HostList(c *gin.Context) {
	appG := app.Gin{C: c}

	var host models.Host
	list, err := host.Get(-1)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, list)
	return
}

func HostInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	hostID, err := strconv.Atoi(c.Param("id"))

	var host models.Host
	list, err := host.Get(hostID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, list)
	return
}

func HostsRemove(c *gin.Context) {
	appG := app.Gin{C: c}

	var arr []int
	removeListOfFail := arr[:]

	json := make(map[string][]int) //注意该结构接受的内容
	c.BindJSON(&json)
	// fmt.Println(json)

	var host models.Host
	for _, v := range json["array"] {
		if err := host.Delete(v); err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}

	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}

func HostUpdate(c *gin.Context) {
	appG := app.Gin{C: c}
	//ctx := context.Background()

	bodyJson := make(map[string]interface{})
	c.BindJSON(&bodyJson)
	host := models.Host{
		Model: models.Model{},
		Name: bodyJson["Name"].(string),
		ViaSocket: bodyJson["ViaSocket"].(bool),
		DockerEngineURL: bodyJson["DockerEngineURL"].(string),
		HostIP: bodyJson["HostIP"].(string),
		TLS: bodyJson["TLS"].(bool),
	}

	if err := host.Update(bodyJson["Id"]); err != nil {
		log.Println(err)
		appG.Response(
			http.StatusInternalServerError,
			e.Error,
			nil,
		)
		return
	}

	appG.Response(
		http.StatusOK,
		e.Success,
		nil,
	)
	return

}
