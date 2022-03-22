package hook

import (
	"autom/service/hook/body"
	"autom/util/request"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/sirupsen/logrus"
)

type GiteeHook struct {
	hook
	token   string
	bodyRaw []byte
}

func (h *GiteeHook) parseBody() {
	var b body.Body
	h.bodyRaw, _ = io.ReadAll(h.c.Request.Body)
	h.c.Request.Body = io.NopCloser(bytes.NewBuffer(h.bodyRaw))
	request.Bind(h.c, &b)
	h.body = &b
}

func (h *GiteeHook) CheckRight(conf *HookContainerConfig) bool {

	token := h.hook.c.GetHeader("X-Gitee-Token")

	if token == "" && conf.Token != "" {
		logrus.Warnf("未接收token")
		return false
	}

	if token != conf.Token {

		timestamp := h.hook.c.GetHeader("X-Gitee-Timestamp")
		h := hmac.New(sha256.New, []byte(conf.Token))
		h.Write([]byte(timestamp + "\n" + conf.Token))

		if token != base64.StdEncoding.EncodeToString(h.Sum(nil)) {
			logrus.Warnf("token不匹配")
			return false
		}
	}

	return h.hook.CheckRight(conf)
}

func (h *GiteeHook) GitType() string {
	return "gitee"
}
