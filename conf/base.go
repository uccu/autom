package conf

import (
	"autom/conf/dev"
	"autom/conf/prod"
	structs "autom/conf/structs"
	"autom/conf/test"
	"io/ioutil"
	"os"
	"time"
)

var Base structs.BaseConf
var Http structs.HttpConf
var Docker structs.Docker
var DebugMode bool
var ConfPath string
var PidPath string
var TimeZone *time.Location
var Log structs.LogConf
var ENV string

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

	Http = Base.Http
	DebugMode = Base.DebugMode
	ConfPath = Base.ConfPath
	PidPath = Base.PidPath
	TimeZone = Base.TimeZone
	Log = Base.Log
	Docker = Base.Docker
}
