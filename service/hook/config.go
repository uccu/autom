package hook

import (
	"autom/conf"
	"autom/service/hook/body"
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

type HookContainerConfig struct {
	Name    string            `json:"name"`   // 项目名字
	Type    string            `json:"type"`   // 推送类型 push/tag_push
	Url     string            `json:"url"`    // git拉取地址
	Branch  *string           `json:"branch"` // 分支，当标签为false有效
	Ip      *string           `json:"ip"`
	Token   string            `json:"token"`
	Volumes map[string]string `json:"volumes"`
	NetWork *HookNetWorkConfig
}

type HookContainerConfigList []*HookContainerConfig

func (m HookContainerConfigList) Filter(b body.Body) HookContainerConfigList {
	list := HookContainerConfigList{}
	for _, v := range m {
		if v.Name == b.GetName() {

			if v.IsTagPsuh() && b.IsTagPsuh() {
				branch := b.GetBranch()
				v.Branch = &branch
				list = append(list, v)
			}

			if v.IsPush() && b.IsPush() && v.Branch != nil && *v.Branch == b.GetBranch() {
				list = append(list, v)
			}
		}
	}
	return list
}

func (c *HookContainerConfig) GetName() string {
	return c.Name
}

func (c *HookContainerConfig) IsTagPsuh() bool {
	return c.Type == "tag_push"
}

func (c *HookContainerConfig) IsPush() bool {
	return c.Type == "push"
}

func (c *HookContainerConfig) GetUrl() string {
	return c.Url
}

func (c *HookContainerConfig) GetBranch() string {
	if c.Branch != nil {
		return *c.Branch
	}
	return ""
}

func (c *HookContainerConfig) GetImageName() string {
	return c.GetName() + ":" + c.GetBranch()
}

func (c *HookContainerConfig) GetNetWorkName() string {
	if c.NetWork.Subnet == nil {
		return "default"
	}
	return c.NetWork.NetWorkName
}

func (c *HookContainerConfig) GetIp() string {
	if c.Ip == nil {
		return ""
	}
	return *c.Ip
}
func (c *HookContainerConfig) GetVolumes() map[string]string {
	return c.Volumes
}

func importConfig() HookContainerConfigList {

	hookConfig := HookConfig{}

	confPath, err := conf.GetConfPath()
	if err != nil {
		logrus.Warnf("配置文件 %s 获取失败！", confPath)
		return nil
	}

	content, err := ioutil.ReadFile(confPath)
	if err != nil {
		logrus.Warnf("配置文件 %s 读取失败！", confPath)
		return nil
	}

	err = json.Unmarshal(content, &hookConfig)
	if err != nil {
		logrus.Warn("配置文件 %s 解析失败！", confPath)
		return nil
	}

	return hookConfig.ToHookContainer()

}
