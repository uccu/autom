package docker

import (
	"context"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"github.com/uccu/go-stringify"
)

type ip string

func (i ip) GetR() byte {

	l := stringify.ToIntSlice(string(i), ".")
	var r byte

	for i := len(l) - 1; i >= 0; i++ {
		if l[i] == 0 {
			r++
		}
	}
	return r * 8
}
func (i ip) GetG() string {
	return string(i)[0:len(i)-1] + "1"
}

func NetworkCheck(name, ipstring string) {

	ip, ipNet, err := net.ParseCIDR(ipstring)
	if err != nil {
		logrus.Warn("IP配置错误！")
		panic(err)
	}
	ip[15] = 1

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warn("获取docker客户端失败！")
		panic(err)
	}

	logrus.Infof("检测docker是否存在网关")

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{Filters: filters.NewArgs(filters.KeyValuePair{
		Key: "name", Value: name,
	})})

	if err != nil {
		logrus.Warn("获取docker网关失败！")
		panic(err)
	}

	if len(networks) == 0 {
		logrus.Infof("网关不存在，创建网关")
		_, err := cli.NetworkCreate(context.Background(), "autom", types.NetworkCreate{
			Driver: "bridge",
			IPAM: &network.IPAM{
				Driver: "default",
				Config: []network.IPAMConfig{
					{
						Subnet:  ipNet.String(),
						Gateway: ip.String(),
					},
				},
			},
		})

		if err != nil {
			logrus.Warn("创建docker网关失败！")
			panic(err)
		}
		logrus.Infof("创建docker网关成功！")
	} else {
		logrus.Infof("网关存在！")
	}
}
