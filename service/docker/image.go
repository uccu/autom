package docker

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type buildConf interface {
	GetName() string
	GetImageName() string
}

func ImageBuild(c buildConf) bool {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warn("获取docker客户端失败！")
		panic(err)
	}

	logrus.Infof("开始创建镜像")

	file, err := os.Open(c.GetName() + "/Dockerfile")
	if err != nil {
		logrus.Warnf("Dockerfile打开失败, %s", err.Error())
		return false
	}

	_, err = cli.ImageBuild(context.Background(), file, types.ImageBuildOptions{
		Tags: []string{c.GetImageName()},
	})

	if err != nil {
		logrus.Warnf("镜像创建失败, %s", err.Error())
		return false
	}

	logrus.Infof("镜像创建成功！")

	return true
}
