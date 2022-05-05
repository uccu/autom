package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uccu/autom/conf"
	"github.com/uccu/autom/lib"
	"github.com/uccu/autom/router"
	"github.com/uccu/autom/service/hook"

	"github.com/sirupsen/logrus"
	"github.com/uccu/go-stringify"
)

func serverStart() error {

	lib.Init()
	logrus.Infof("start autom success, port:%d", conf.Http.Port)
	defer lib.Destroy()

	pid := stringify.ToString(os.Getpid())
	pidPath, err := conf.GetPidPath()
	if err != nil {
		return err
	}
	ioutil.WriteFile(pidPath, []byte(pid), 0600)

	quit := make(chan os.Signal)

	go func() {

		for {
			line, err := ioutil.ReadFile(pidPath)
			if string(line) == pid {
				time.Sleep(500 * time.Millisecond)
				continue
			}

			if err == nil {
				ioutil.WriteFile(pidPath, []byte("-1"), 0600)
			}

			quit <- syscall.SIGQUIT
			break
		}
	}()

	router.HttpServerRun()

	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func serverStop() error {

	pidPath, err := conf.GetPidPath()
	if err != nil {
		return err
	}

	ioutil.WriteFile(pidPath, []byte("0"), 0600)

	var times int
	for {
		line, err := ioutil.ReadFile(pidPath)
		if err != nil {
			return err
		}
		if string(line) != "-1" && times < 10 {
			time.Sleep(500 * time.Millisecond)
			times++
			continue
		}
		break
	}

	os.Remove(pidPath)
	fmt.Println("stop autom success")
	return nil

}

func runSingle(id string) error {

	hookContainerConfigList := hook.ImportConfig()

	nameSlice := stringify.ToStringSlice(id, ":")
	if len(nameSlice) < 2 {
		logrus.Warnf("请指定版本: %s！", id)
		return nil
	}
	id = nameSlice[0]
	branch := nameSlice[1]

	hookContainerConfig := hookContainerConfigList.GetById(id)
	if hookContainerConfig == nil {
		logrus.Warnf("配置ID %s 获取失败或不存在！", id)
		return nil
	}
	hookContainerConfig.Branch = &branch

	hook.RunPush(hookContainerConfig)
	return nil
}
