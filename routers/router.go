package routers

import (
	"gitee.com/favefan/doconsole/pkg/setting"
	v1 "gitee.com/favefan/doconsole/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter init router, return *gin.Engine.
func InitRouter() *gin.Engine {
    gin.SetMode(setting.ServerSetting.RunMode)

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    apiv1 := r.Group("/api/v1")
    {
        apiv1.GET("/containers", v1.ListContainers)
        apiv1.POST("/containers/create", v1.CreateContainer)
    }

    return r
}