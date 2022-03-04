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
	DebugMode bool
	TimeZone  *time.Location
	ConfPath  string
	PidPath   string
	WorkDir   string
	Http      HttpConf
	Log       LogConf
	Docker    Docker
}
