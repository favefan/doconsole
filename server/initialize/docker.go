package initialize

import (
	"fmt"
	"github.com/docker/docker/client"
)

const dockerClientVersion = "1.41"

type Host struct {
	Name string `form:"Name"`
	ViaSocket bool `form:"ViaSocket"`
	DockerEngineURL string `form:"URL"`
	HostIP string `form:"HostIP"`
	TLS bool `form:"TLS"`
	Default bool `form:"Default"`
}

func DockerLocalClientSetup () (*client.Client, error) {
	return client.NewClientWithOpts(client.WithVersion(dockerClientVersion))
}

func DockerTCPClientSetup (host string) (*client.Client, error) {
	return client.NewClientWithOpts(client.WithHost(host), client.WithVersion(dockerClientVersion))
}

func CreateDockerClient (host *Host) (*client.Client, error) {
	fmt.Println(host)

	if host.ViaSocket == true {
		return DockerLocalClientSetup()
	} else {
		return DockerTCPClientSetup(host.DockerEngineURL)
	}
}
