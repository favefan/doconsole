package routers

import (
	"gitee.com/favefan/doconsole/initialize"
	v1 "gitee.com/favefan/doconsole/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter init router, return *gin.Engine.
func InitRouter() *gin.Engine {
	gin.SetMode(initialize.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//apiAuth := r.Group("/auth")
	//{
	//	apiAuth.POST("/login", api.Login)
	//	apiAuth.GET("/logout", api.Logout)
	//}
	// r.GET("/ws/:id", v1.ContainerExecAttach)

	apiV1 := r.Group("/api/v1")
	// apiV1.Use(jwt.JWT())
	{
		// docker engine
		{
			apiV1.GET("/docker/endpoint", v1.DaemonHost)
			apiV1.GET("/docker/version", v1.ClientVersion)
			apiV1.PUT("/docker/client", v1.CreateClient)
		}
		// containers
		{
			apiV1.GET("/containers/json", v1.ContainerList)
			apiV1.GET("/containers/json/:id", v1.ContainerInspect)
			apiV1.POST("/containers/update/:id", v1.ContainerUpdate)
			apiV1.GET("/containers/start", v1.ContainerStart)
			apiV1.GET("/containers/stop", v1.ContainerStop)
			apiV1.GET("/containers/restart", v1.ContainerRestart)
			apiV1.GET("/containers/pause", v1.ContainerPause)
			apiV1.DELETE("/containers/:id", v1.ContainerRemove)
			apiV1.POST("/containers/create", v1.ContainerCreate)
			apiV1.POST("/containers/exec/:id", v1.ContainerExecCreate)
			apiV1.GET("/containers/exec/:id", v1.ContainerExecAttach)
			apiV1.GET("/containers/logs/:id", v1.ContainerLogs)
			apiV1.GET("/containers/stats/:id", v1.ContainerStats)
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
			apiV1.POST("/networks/connect/:id", v1.NetworkConnect)
			apiV1.POST("/networks/disconnect/:id", v1.NetworkDisconnect)
			apiV1.DELETE("/networks", v1.NetworksRemove)
		}
		// volumes
		{
			apiV1.GET("/volumes/json", v1.VolumeList)
			apiV1.GET("/volumes/json/:name", v1.VolumeInspect)
			apiV1.POST("/volumes/create", v1.VolumeCreate)
			apiV1.DELETE("/volumes", v1.VolumesRemove)
		}
		// registries
		{
			apiV1.POST("/registries/create", v1.RegistryCreate)
			apiV1.PUT("/registries", v1.RegistryUpdate)
			apiV1.GET("/registries/json", v1.RegistryList)
			apiV1.GET("/registries/json/:id", v1.RegistryInspect)
			apiV1.DELETE("/registries", v1.RegistriesRemove)
		}
		// registries
		{
			apiV1.POST("/hosts/create", v1.HostCreate)
			apiV1.PUT("/hosts", v1.HostUpdate)
			apiV1.GET("/hosts/json", v1.HostList)
			apiV1.GET("/hosts/json/:id", v1.HostInspect)
			apiV1.DELETE("/hosts", v1.HostsRemove)
		}
	}

	return r
}
