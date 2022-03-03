package structs

import (
	"time"

	"github.com/sirupsen/logrus"
)

type HttpConf struct {
	Port           int
	HeaderCheck    bool
	TrustedProxies []string
}

type LogConf struct {
	Level      logrus.Level
	TimeLayout string
	Path       string
}

type Docker struct {
	DockerApiVersion string
}

type BaseConf struct {
	Http      HttpConf
	TimeZone  *time.Location
	DebugMode bool
	Log       LogConf
	ConfPath  string
	PidPath   string
	Docker    Docker
}
