package git

import (
	"os/exec"

	"github.com/uccu/autom/util/fs"

	"github.com/sirupsen/logrus"
)

type gitConf interface {
	GetUrl() string
	GetBranch() string
	GetName() string
}

func Clone(c gitConf) bool {

	cmd := exec.Command("git", "clone", "-b", c.GetBranch(), c.GetUrl(), c.GetName()+"/resp")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git clone failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git clone success")
	return true
}

func HasCache(c gitConf) bool {
	return fs.PathExists(c.GetName() + "/resp")
}

func Fetch(c gitConf) bool {

	cmd := exec.Command("git", "fetch")
	cmd.Dir = c.GetName() + "/resp"

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git fetch failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git fetch success")
	return true
}

func Checkout(c gitConf) bool {

	cmd := exec.Command("git", "checkout", c.GetBranch())
	cmd.Dir = c.GetName() + "/resp"

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git checkout failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git checkout success")
	return true
}

func Pull(c gitConf) bool {

	cmd := exec.Command("git", "pull")
	cmd.Dir = c.GetName() + "/resp"

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git pull failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git pull success")
	return true
}

func Archive(c gitConf) bool {

	if !fs.PathExists(c.GetName() + "/resp/Dockerfile") {
		logrus.Infof("repository's Dockerfile not exist, non archive")
		return false
	}

	cmd := exec.Command("git", "archive", c.GetBranch(), "-o", "../"+c.GetName()+".tar")
	cmd.Dir = c.GetName() + "/resp"

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Warnf("repository git archive failed: %s\n%s", err.Error(), output)
		logrus.Warnf(err.Error())
		return false
	}

	logrus.Infof("repository git archive success")
	return true
}
