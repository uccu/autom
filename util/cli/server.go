package cli

import (
	"autom/conf"
	"autom/lib"
	"autom/router"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uccu/go-stringify"
)

func serverStart() error {

	lib.Init()
	logrus.Infof("start autom success, port:%d", conf.Http.Port)
	defer lib.Destroy()

	pid := stringify.ToString(os.Getpid())
	pidPath := conf.PidPath

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

	pidPath := conf.PidPath

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
