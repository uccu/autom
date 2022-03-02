package lib

import (
	"autom/conf"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type writers []io.Writer

func newWriters() *writers {
	ws := make(writers, 0)
	return &ws
}

func (ws *writers) WithFileWriter() *writers {

	if conf.Log.Path == "" {
		return ws
	}

	p := path.Dir(conf.Log.Path)
	if p == "." {
		return ws
	}

	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		fmt.Printf("create log path dir `%s` err: %s", p, err.Error())
		return ws
	}

	file, err := os.OpenFile(conf.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("open log file `%s` err: %s", conf.Log.Path, err.Error())
		return ws
	}

	*ws = append(*ws, file)

	return ws
}

func (ws *writers) WithStdWriter() *writers {
	*ws = append(*ws, os.Stdout)
	return ws
}

//模块初始化
func Init() {

	os.Setenv("DOCKER_API_VERSION", conf.Docker.DockerApiVersion)

	time.Local = conf.TimeZone
	color.NoColor = false

	logger := logrus.StandardLogger()
	logger.SetOutput(io.MultiWriter(*newWriters().WithFileWriter().WithStdWriter()...))

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: conf.Log.TimeLayout,
	})
	logrus.SetLevel(conf.Log.Level)

}

//公共销毁函数
func Destroy() {
	logrus.Info("stop autom success")
}
