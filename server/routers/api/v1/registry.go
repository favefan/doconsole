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

func RegistryCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	//ctx := context.Background()

	bodyJson := make(map[string]interface{})
	c.BindJSON(&bodyJson)
	registry := models.Registry{
		Model:    models.Model{},
		Name:     bodyJson["Name"].(string),
		URL:      bodyJson["URL"].(string),
		NeedAuth: bodyJson["NeedAuth"].(bool),
		Username: bodyJson["Username"].(string),
		Password: bodyJson["Password"].(string),
		Comment:  "",
	}

	if err := registry.Create(); err != nil {
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

func RegistryList(c *gin.Context) {
	appG := app.Gin{C: c}

	var registry models.Registry
	list, err := registry.Get(-1)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, list)
	return
}

func RegistryInspect(c *gin.Context) {
	appG := app.Gin{C: c}
	registryID, err := strconv.Atoi(c.Param("id"))

	var registry models.Registry
	list, err := registry.Get(registryID)
	if err != nil {
		log.Println(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, list)
	return
}

func RegistriesRemove(c *gin.Context) {
	appG := app.Gin{C: c}

	var arr []int
	removeListOfFail := arr[:]

	json := make(map[string][]int) //注意该结构接受的内容
	c.BindJSON(&json)
	// fmt.Println(json)

	var registry models.Registry
	for _, v := range json["array"] {
		if err := registry.Delete(v); err != nil {
			log.Println(err)
			removeListOfFail = append(removeListOfFail, v)
		}
	}

	appG.Response(http.StatusOK, e.Success, removeListOfFail)
	return
}
