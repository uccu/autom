package git

import (
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type gitConf interface {
	GetUrl() string
	GetBranch() string
	GetName() string
}

func Clone(c gitConf) bool {

	cmd := exec.Command("git", "clone", "-b", c.GetBranch(), "--single-branch", c.GetUrl(), c.GetName())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git clone failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git clone success!")
	return true
}

func Remove(c gitConf) {
	err := os.RemoveAll(c.GetName())
	if err != nil {
		logrus.Infof("repository git remove failed!")
		return
	}
	logrus.Infof("repository git remove success!")
}
