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

	//apiAuth := r.Group("/auth")
	//{
	//	apiAuth.POST("/login", api.Login)
	//	apiAuth.GET("/logout", api.Logout)
	//}

	apiV1 := r.Group("/api/v1")
	// apiV1.Use(jwt.JWT())
	{
		// containers
		{
			apiV1.GET("/containers/json", v1.ContainerList)
			apiV1.GET("/containers/json/:id", v1.ContainerInspect)
			apiV1.GET("/containers/start", v1.ContainerStart)
			apiV1.GET("/containers/stop", v1.ContainerStop)
			apiV1.GET("/containers/restart", v1.ContainerRestart)
			apiV1.GET("/containers/pause", v1.ContainerPause)
			apiV1.DELETE("/containers/:id", v1.ContainerRemove)
			apiV1.POST("/containers/create", v1.CreateContainer)
		}
		// images
		{
			apiV1.GET("/images/json", v1.ImageList)
			apiV1.GET("/images/search", v1.ImageSearch)
			apiV1.GET("/images/json/:id", v1.ImageInspect)
			apiV1.GET("/images/history/:id", v1.ImageHistory)
			apiV1.GET("/images/pull", v1.ImagePull)
			apiV1.DELETE("/images/:id", v1.ImageRemove)
			apiV1.GET("/distribution/json/:id", v1.ImageInspectFromRegistry)
		}
		// networks
		{
			apiV1.GET("/networks/json", v1.NetworkList)
			apiV1.GET("/networks/json/:id", v1.NetworkInspect)
			apiV1.POST("/networks/create", v1.NetworkCreate)
			apiV1.DELETE("/networks", v1.NetworksRemove)
		}
		// volumes
		{
			apiV1.GET("/volumes/json", v1.VolumeList)
		}
	}

	return r
}
