package main_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestExec2(t *testing.T) {

	content, _ := ioutil.ReadFile("a.json")

	var a interface{}
	json.Unmarshal(content, &a)
	b, _ := json.Marshal(a)

	fmt.Println(string(b))

	h := hmac.New(sha256.New, []byte("affe63ae92f9f8b451b0fb3979cf9dbc5ed52b2b"))
	h.Write(b)
	s := h.Sum(nil)
	sig := hex.EncodeToString(s)

	fmt.Println(sig)

	// b890b2353e5f64e0a71ad85ddf1077d2f28915edff1684eaf0b431b464b32ae0
	// 2707a7bd9eaa6c49c931c3eff7406a8568e24dded8f796854245d80ae85e106b
}
