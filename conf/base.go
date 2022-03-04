package conf

import (
	"autom/conf/dev"
	"autom/conf/prod"
	structs "autom/conf/structs"
	"autom/conf/test"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var Base structs.BaseConf
var Http structs.HttpConf
var Docker structs.Docker
var TimeZone *time.Location
var Log structs.LogConf
var ENV string
var DebugMode bool

func setEnvironment() {
	content, _ := ioutil.ReadFile("environment")
	env := string(content)
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}

	if env != "dev" && env != "test" {
		env = "prod"
	}

	ENV = env
}

func init() {

	setEnvironment()
	switch ENV {
	case "dev":
		Base = dev.Base
	case "test":
		Base = test.Base
	default:
		Base = prod.Base
	}

	Base.WorkDir, _ = os.Getwd()

	Http = Base.Http
	TimeZone = Base.TimeZone
	Log = Base.Log
	Docker = Base.Docker
	DebugMode = Base.DebugMode
}

func GetPidPath() (string, error) {
	pidPath := Base.PidPath
	if !path.IsAbs(Base.PidPath) {
		pidPath = path.Join(Base.WorkDir, Base.PidPath)
	}
	p := path.Dir(pidPath)
	if p != "." {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return pidPath, err
		}
	}
	return pidPath, nil
}

func GetConfPath() (string, error) {
	confPath := Base.ConfPath
	if !path.IsAbs(Base.ConfPath) {
		confPath = path.Join(Base.WorkDir, Base.ConfPath)
	}
	p := path.Dir(confPath)
	if p != "." {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return confPath, err
		}
	}
	return confPath, nil
}

func GetLogPath() (string, error) {

	logPath := Base.Log.Path
	if logPath == "" {
		return "", nil
	}

	if !path.IsAbs(Base.Log.Path) {
		logPath = path.Join(Base.WorkDir, Base.Log.Path)
	}
	p := path.Dir(logPath)
	if p != "." {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			fmt.Printf("create log path dir `%s` err: %s", p, err.Error())
			return logPath, err
		}
	}
	return logPath, nil
}
