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

	err := cmd.Run()
	if err != nil {
		logrus.Warnf(err.Error())
		return false
	}

	return true
}

func Remove(c gitConf) {
	err := os.RemoveAll(c.GetName())
	if err != nil {
		panic(err)
	}
}
