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
	IsTag() bool
	GetBranch() string
}

func ContainerCreate(c containerConf) bool {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warnf("get docker failed: %s", err.Error())
		return false
	}

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

	name := c.GetName()
	if !c.IsTag() {
		name += "_" + c.GetBranch()
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", c.GetName())),
	})
	if err != nil {
		logrus.Warnf("get docker container list failed: %s", err.Error())
		return false
	}

	if len(containers) > 0 {
		container := containers[0]

		if container.State == "running" {
			err := cli.ContainerStop(context.Background(), container.ID, nil)
			if err != nil {
				logrus.Warnf("stop docker old container failed: %s", err.Error())
				return false
			}

		}

		err := cli.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{})
		if err != nil {
			logrus.Warnf("remove docker old container failed: %s", err.Error())
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
		logrus.Warnf("craete docker new container failed: %s", err.Error())
		return false
	}

	logrus.Infof("create docker new container success, ID: %s", b.ID)

	err = cli.ContainerStart(context.Background(), b.ID, types.ContainerStartOptions{})
	if err != nil {
		logrus.Warnf("start docker container failed: %s", err.Error())
		return false
	}
	logrus.Infof("start docker container success, ID: %s", b.ID)

	return true
}
