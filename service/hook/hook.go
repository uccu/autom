package hook

import (
	"autom/service/docker"
	"autom/util/request"

	"autom/service/git"

	"github.com/gin-gonic/gin"
)

type Hook interface {

	// gitlab/github/git/gitee等
	GitType() string

	// Push Hook/Tag Push Hoo
	Event() string

	CheckRight() bool

	ParseBody()

	GetBranch() string
	GetName() string
	Run() bool
}

func NewHookClient(c *gin.Context) Hook {

	h := hook{
		c:       c,
		gitType: "gitlab",
	}

	h.ParseBody()
	h.conf = importConfig().Get(h.GetName(), h.Event() == "tag_push")

	if h.conf == nil {
		return nil
	}

	if c.GetHeader("X-Gitlab-Event") != "" {
		return &GitlabHook{
			hook: h,
		}
	}
	return &h
}

type hook struct {
	c       *gin.Context
	gitType string
	event   string
	body    *HookBody
	conf    *HookContainerConfig
}

func (h *hook) GitType() string {
	return h.gitType
}

func (h *hook) Event() string {
	return h.event
}

func (h *hook) CheckRight() bool {
	if h.body.CheckoutSha == nil {
		return false
	}
	return true
}

func (h *hook) ParseBody() {
	var b HookBody
	request.Bind(h.c, &b)
	h.body = &b
	h.event = h.body.GetEvent()
}

func (h *hook) GetBranch() string {
	return h.body.GetBranch()
}

func (h *hook) GetName() string {
	return h.body.GetName()
}

func (h *hook) GetUrl() string {
	return h.conf.Url
}

func (h *hook) GetImageName() string {
	return h.GetName() + ":" + h.GetBranch()
}

func (h *hook) GetNetWorkName() string {
	if h.conf.NetWork.Subnet == nil {
		return "default"
	}
	return h.conf.NetWork.NetWorkName
}

func (h *hook) GetIp() string {
	if h.conf.Ip == nil {
		return ""
	}
	return *h.conf.Ip
}
func (h *hook) GetVolumes() map[string]string {
	return h.conf.Volumes
}

func (h *hook) Run() bool {

	if !h.conf.Tag {
		if h.conf.Branch == nil || *h.conf.Branch != h.GetBranch() {
			return false
		}
	}

	if h.conf.NetWork.Subnet != nil {
		docker.NetworkCheck(h.conf.NetWork.NetWorkName, *h.conf.NetWork.Subnet)
	}

	if !git.Clone(h) {
		return false
	}
	defer git.Remove(h)

	if !docker.ImageBuild(h) {
		return false
	}
	if !docker.ContainerCreate(h) {
		return false
	}

	return true

}
