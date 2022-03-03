package hook

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type HookConfig []*HookNetWorkConfig

func (hookConfig HookConfig) ToHookContainer() HookContainerConfigList {
	hookContainer := HookContainerConfigList{}
	for _, hookNetWorkConfig := range hookConfig {
		for _, hookContainerConfig := range hookNetWorkConfig.Containers {
			hookContainerConfig.NetWork = hookNetWorkConfig
		}
		hookContainer = append(hookContainer, hookNetWorkConfig.Containers...)
	}
	return hookContainer
}

type HookNetWorkConfig struct {
	NetWorkName string                  `json:"network_name"`
	Subnet      *string                 `json:"subnet"`
	Containers  HookContainerConfigList `json:"containers"`
}

type HookContainerConfigList []*HookContainerConfig

func (m HookContainerConfigList) Get(name string, isTag bool) *HookContainerConfig {
	for _, v := range m {
		if v.Name == name && v.Tag == isTag {
			return v
		}
	}
	return nil
}

type HookContainerConfig struct {
	Name    string            `json:"name"`   // 项目名字
	Tag     bool              `json:"tag"`    // 是否是标签
	Url     string            `json:"url"`    // git拉取地址
	Branch  *string           `json:"branch"` // 分支，当标签为false有效
	Ip      *string           `json:"ip"`
	Token   string            `json:"token"`
	Volumes map[string]string `json:"volumes"`
	NetWork *HookNetWorkConfig
}

func importConfig() HookContainerConfigList {

	hookConfig := HookConfig{}

	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		logrus.Warn("config.json 读取失败！")
		panic(err)
	}

	err = json.Unmarshal(content, &hookConfig)
	if err != nil {
		logrus.Warn("config.json 读取失败！")
		panic(err)
	}

	return hookConfig.ToHookContainer()

}
