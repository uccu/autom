package hook

import (
	"github.com/uccu/autom/service/docker"
	"github.com/uccu/autom/service/hook/body"
	"github.com/uccu/autom/util/request"

	"github.com/uccu/autom/service/git"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Body interface {
	GetName() string
	IsTagPsuh() bool
	IsPush() bool
	IsInvalid() bool
	GetBranch() string
}

type Hook interface {
	getConfs() HookContainerConfigList
	parseBody()
	parseConfs()

	// gitlab/github/git/gitee等
	GitType() string

	CheckRight(*HookContainerConfig) bool
}

func NewHookClient(c *gin.Context) Hook {
	h := parseGitType(c, &hook{c: c})
	h.parseBody()
	h.parseConfs()
	return h
}

type hook struct {
	c    *gin.Context
	body Body
	conf HookContainerConfigList
}

func (h *hook) GitType() string {
	return ""
}

func (h *hook) CheckRight(conf *HookContainerConfig) bool {
	return h.body.IsInvalid()
}

func (h *hook) parseBody() {
	var b body.Body
	request.Bind(h.c, &b)
	h.body = &b
}

func (h *hook) parseConfs() {
	if h.body != nil {
		confs := ImportConfig()
		h.conf = confs.Filter(h.body)
	}
}

func (h *hook) getConfs() HookContainerConfigList {
	return h.conf
}

func Run(h Hook) {
	for _, conf := range h.getConfs() {
		if h.CheckRight(conf) {
			RunPush(conf)
		}
	}
}

func RunPush(conf *HookContainerConfig) bool {

	if conf.NetWork.Subnet != nil {
		docker.NetworkCheck(conf.NetWork.NetWorkName, *conf.NetWork.Subnet)
	}

	if git.HasCache(conf) {
		logrus.Infof("repository git cache exist")
		if !git.Fetch(conf) {
			return false
		}

		if conf.IsPush() {
			if !git.Checkout(conf) {
				return false
			}

			if !git.Pull(conf) {
				return false
			}
		}
	} else {
		logrus.Infof("repository git cache not exist")
		if !git.Clone(conf) {
			return false
		}
	}

	if !git.Archive(conf) {
		return false
	}

	if !docker.ImageBuild(conf) {
		return false
	}
	if !docker.ContainerCreate(conf) {
		return false
	}

	return true
}

func parseGitType(c *gin.Context, h *hook) Hook {

	if c.GetHeader("X-Gitlab-Event") != "" {
		return &GitlabHook{
			hook: *h,
		}
	}

	if c.GetHeader("X-Gitee-Event") != "" {
		return &GiteeHook{
			hook: *h,
		}
	}

	if c.GetHeader("X-Github-Event") == "push" {
		return &GithubHook{
			hook: *h,
		}
	}

	return h
}
