package hook

import (
	"autom/conf"
	"autom/http_error"
)

type GitlabHook struct {
	hook
	token string
}

func (h *GitlabHook) CheckRight() bool {

	token := h.hook.c.GetHeader("X-Gitlab-Token")

	if token == "" && conf.Http.Token != "" {
		panic(http_error.NoXGitlabToken)
	}

	if token != conf.Http.Token {
		panic(http_error.XGitlabTokenNotMatch)
	}

	return h.hook.CheckRight()
}
