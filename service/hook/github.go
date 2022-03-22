package hook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"github.com/uccu/autom/service/hook/body"
	"github.com/uccu/autom/util/request"

	"github.com/sirupsen/logrus"
)

type GithubHook struct {
	hook
	token   string
	bodyRaw []byte
}

func (h *GithubHook) parseBody() {
	var b body.Body
	h.bodyRaw, _ = io.ReadAll(h.c.Request.Body)
	h.c.Request.Body = io.NopCloser(bytes.NewBuffer(h.bodyRaw))
	request.Bind(h.c, &b)
	h.body = &b
}

func (h *GithubHook) CheckRight(conf *HookContainerConfig) bool {

	token := h.hook.c.GetHeader("X-Hub-Signature-256")

	if token == "" && conf.Token != "" {
		logrus.Warnf("未接收token")
		return false
	}

	hm := hmac.New(sha256.New, []byte(conf.Token))
	hm.Write(h.bodyRaw)
	sig := hex.EncodeToString(hm.Sum(nil))

	if token != "sha256="+string(sig) {
		logrus.Warnf("token不匹配")
		return false
	}

	return h.hook.CheckRight(conf)
}

func (h *GithubHook) GitType() string {
	return "github"
}
