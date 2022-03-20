package docker

import (
	"context"
	"net"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"github.com/uccu/go-stringify"
)

type ip string

var networkChecklock sync.Mutex

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

func NetworkCheck(name, ipstring string) bool {

	networkChecklock.Lock()
	defer networkChecklock.Unlock()

	ip, ipNet, err := net.ParseCIDR(ipstring)
	if err != nil {
		logrus.Warnf("ipstring %s wrong: %s", ipstring, err.Error())
		return false
	}
	ip[15] = 1

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logrus.Warnf("get docker failed: %s", err.Error())
		return false
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{Filters: filters.NewArgs(filters.KeyValuePair{
		Key: "name", Value: name,
	})})

	if err != nil {
		logrus.Warnf("get docker network list failed: %s", err.Error())
		return false
	}

	if len(networks) == 0 {
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
			logrus.Warnf("create docker network failed: %s", err.Error())
			return false
		}
		logrus.Infof("create docker network success")
	} else {
		logrus.Infof("docker network exist")
	}

	return true
}
