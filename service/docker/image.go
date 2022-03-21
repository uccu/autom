package docker

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"

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
		logrus.Warnf("get docker failed: %s", err.Error())
		return false
	}

	tar, err := ioutil.ReadFile(c.GetName() + "/" + c.GetName() + ".tar")
	if err != nil {
		logrus.Warnf("read tar file failed: %s", err.Error())
		return false
	}

	res, err := cli.ImageBuild(context.Background(), bytes.NewBuffer(tar), types.ImageBuildOptions{
		Tags: []string{c.GetImageName()},
	})

	if err != nil {
		logrus.Warnf("docker image build failed: %s", err.Error())
		return false
	}

	io.ReadAll(res.Body)

	logrus.Infof("docker image build success")
	return true
}
