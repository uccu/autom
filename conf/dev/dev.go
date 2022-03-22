package dev

import (
	"time"

	. "github.com/uccu/autom/conf/structs"

	"github.com/sirupsen/logrus"
)

var Base BaseConf = BaseConf{
	DebugMode: true,
	PidPath:   "/opt/autom/autom.pid",
	ConfPath:  "config.json",
	TimeZone:  time.FixedZone("CST", 8*3600),
	Http: HttpConf{
		Port:        2333,
		HeaderCheck: true,
	},

	Log: LogConf{
		Level:      logrus.TraceLevel,
		TimeLayout: "2006/01/02 - 15:04:05.000",
	},
	Docker: Docker{
		DockerApiVersion: "1.41",
	},
}
