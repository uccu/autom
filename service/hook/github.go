package hook

import (
	"autom/service/hook/body"
	"autom/util/request"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type GithubHook struct {
	hook
	token   string
	bodyRaw []byte
}

func (h *GithubHook) parseBody() {
	var b body.GithubBody
	h.bodyRaw, _ = io.ReadAll(h.c.Request.Body)
	h.c.Request.Body = io.NopCloser(bytes.NewBuffer(h.bodyRaw))
	request.Bind(h.c, &b)
	h.body = &b
}

func (h *GithubHook) CheckRight(conf *HookContainerConfig) bool {

	hm := hmac.New(sha256.New, []byte(conf.Token))
	hm.Write(h.bodyRaw)
	sig := hex.EncodeToString(hm.Sum(nil))

	token := h.hook.c.GetHeader("X-Hub-Signature-256")

	fmt.Println(h.hook.c.Request.Header)

	if token == "" && conf.Token != "" {
		logrus.Warnf("未接收token")
		return false
	}

	if token != "sha256="+string(sig) {
		logrus.Warnf("token不匹配")
		return false
	}

	return h.hook.CheckRight(conf)
}

func (h *GithubHook) GitType() string {
	return "github"
}
