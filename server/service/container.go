package service

import "fmt"

type ContainerCreationBody struct {
	Name string `form:"Name"`
	Image string `form:"Image"`
	ExposedPorts map[string]interface{} `form:"ExposedPorts"`
	Env []string `form:"Env"`
	Cmd []string `form:"Cmd"`
	Entrypoint string `form:"Entrypoint"`
	WorkingDir string `form:"WorkingDir"`
	User string `form:"User"`
	Hostname string `form:"Hostname"`
	Domainname string `form:"Domainname"`
	MacAddress string `form:"MacAddress"`
	Tty bool `form:"Tty"`
	OpenStdin bool `form:"OpenStdin"`
	HostConfig struct{
		Binds []string `form:"HostConfig.Binds"`
		PortBindings interface{} `form:"HostConfig.PortBindings"`
		PublishAllPorts bool `form:"HostConfig.PublishAllPorts"`
		AutoRemove bool `form:"HostConfig.AutoRemove"`
		RestartPolicy struct {
			Name string `form:"HostConfig.RestartPolicy.Name"`
		}
	}
	NetworkMode string `form:"NetworkMode"`
	Volumes interface{} `form:"Volumes"`
	NetworkingConfig struct {
		EndpointsConfig interface{} `form:"NetworkingConfig.EndpointsConfig"`
	}
}

func ContainerCreate(body ContainerCreationBody) (string, error) {
	fmt.Println(&body)
	
	return "ok", nil
}
