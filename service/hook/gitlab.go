package hook

import (
	"autom/service/hook/body"
	"autom/util/request"

	"github.com/sirupsen/logrus"
)

type GitlabHook struct {
	hook
	token string
}

func (h *GitlabHook) parseBody() {
	var b body.GitlabBody
	request.Bind(h.c, &b)
	h.body = &b
}

func (h *GitlabHook) CheckRight(conf *HookContainerConfig) bool {

	token := h.hook.c.GetHeader("X-Gitlab-Token")

	if token == "" && conf.Token != "" {
		logrus.Warnf("未接收token")
		return false
	}

	if token != conf.Token {
		logrus.Warnf("token不匹配")
		return false
	}

	return h.hook.CheckRight(conf)
}

func (h *GitlabHook) GitType() string {
	return "gitlab"
}
