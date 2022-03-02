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

	log, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("git log: %s", string(log))
		return false
	}

	logrus.Infof("git log: %s", string(log))
	return true
}

func Remove(c gitConf) {
	err := os.RemoveAll(c.GetName())
	if err != nil {
		panic(err)
	}
}
