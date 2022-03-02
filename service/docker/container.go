package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type containerConf interface {
	GetName() string
	GetImageName() string
	GetIp() string
	GetNetWorkName() string
	GetVolumes() map[string]string
}

func ContainerCreate(c containerConf) string {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warn("获取docker客户端失败！")
		panic(err)
	}

	m := []mount.Mount{}
	for k, v := range c.GetVolumes() {
		m = append(m, mount.Mount{Type: mount.TypeBind, Source: k, Target: v})
	}

	var nc *network.NetworkingConfig = nil
	if c.GetIp() != "" && c.GetNetWorkName() != "" {
		nc = &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				c.GetNetWorkName(): {
					IPAMConfig: &network.EndpointIPAMConfig{
						IPv4Address: c.GetIp(),
					},
				},
			},
		}
	}

	b, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: c.GetImageName(),
	}, &container.HostConfig{
		Mounts: m,
	}, nc, nil, c.GetName())
	if err != nil {
		logrus.Warn("运行容器失败！")
		panic(err)
	}

	logrus.Infof("运行容器成功, ID: %s", b.ID)

	return b.ID
}
