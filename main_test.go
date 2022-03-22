package main_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFile(t *testing.T) {

	p := "log/warn/2022-02-26.log"

	if p := path.Dir(p); p != "." {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
	}

	_, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	fmt.Println(err)

}
func TestIp(t *testing.T) {

	ip, ipNet, err := net.ParseCIDR("172.18.0.0/16")
	if err != nil {
		logrus.Warn("IP配置错误")
		panic(err)
	}

	ip[15] = 1

	log.Println(ip.String())
	log.Println(ipNet.String())
}

func TestExec(t *testing.T) {
	cmd := exec.Command("git clone git@github.com:uccu/swkoa-config.git -b v0.1.2 --single-branch --no-tags resp")
	whoami, err := cmd.Output()
	fmt.Println(string(whoami))
	fmt.Println(err)

}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func TestSign(t *testing.T) {

	key := "1234567890"
	val := "1647921033177" + "\n" + "1234567890"
	h := hmac.New(sha256.New, []byte(key))

	h.Write([]byte(val))
	b := h.Sum(nil)

	v := base64.StdEncoding.EncodeToString(b)

	fmt.Println(v)

}
