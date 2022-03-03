package docker

import (
	"os"
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

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		logrus.Warnf(err.Error())
		return false
	}

	return true
}
