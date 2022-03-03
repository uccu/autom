package docker

import (
	"context"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
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

func ContainerCreate(c containerConf) bool {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warn("获取docker客户端失败, %s", err.Error())
		return false
	}

	logrus.Infof("创建容器")

	dir, _ := os.Getwd()

	m := []mount.Mount{}
	for k, v := range c.GetVolumes() {

		source := k
		if !path.IsAbs(k) {
			source = path.Join(dir, k)
		}
		err := os.MkdirAll(source, os.ModePerm)
		if err != nil {
			logrus.Warn("path %s err: %s", source, err.Error())
			return false
		}
		m = append(m, mount.Mount{Type: mount.TypeBind, Source: source, Target: v})
	}

	var nc *network.NetworkingConfig = nil
	if c.GetIp() != "" && c.GetNetWorkName() != "default" {
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

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", c.GetName())),
	})
	if err != nil {
		logrus.Warn("容器列表获取失败, %s", err.Error())
		return false
	}

	if len(containers) > 0 {
		container := containers[0]

		if container.State == "running" {
			err := cli.ContainerStop(context.Background(), container.ID, nil)
			if err != nil {
				logrus.Warn("暂停老版本容器失败, %s", err.Error())
				return false
			}

		}

		err := cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{})
		if err != nil {
			logrus.Warn("移除老版本容器失败, %s", err.Error())
			return false
		}
	}

	b, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: c.GetImageName(),
	}, &container.HostConfig{
		NetworkMode: container.NetworkMode(c.GetNetWorkName()),
		Mounts:      m,
	}, nc, nil, c.GetName())
	if err != nil {
		logrus.Warn("创建容器失败, %s", err.Error())
		return false
	}

	logrus.Infof("创建容器成功, ID: %s", b.ID)

	logrus.Infof("启动容器, ID: %s", b.ID)
	err = cli.ContainerStart(context.Background(), b.ID, types.ContainerStartOptions{})
	if err != nil {
		logrus.Warn("启动容器失败, %s", err.Error())
		return false
	}
	logrus.Infof("启动容器成功, ID: %s", b.ID)

	return true
}
