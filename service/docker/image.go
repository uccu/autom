package docker

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

type buildConf interface {
	GetName() string
	GetImageName() string
}

func ImageBuild(c buildConf) bool {

	cmd := exec.Command("docker", "build", "-t", c.GetImageName(), ".")
	cmd.Dir = c.GetName()

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("docker image build failed: %s\n%s", err.Error(), output)
		return false
	}

	logrus.Infof("docker image build success!")
	return true
}
