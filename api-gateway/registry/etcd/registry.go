package etcd

import (
	"github.com/LCY2013/paas/api-gateway/config"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"time"
)

func NewRegistry(config *config.Config) registry.Registry {
	var etcdRegistry = etcd.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = config.RegisterAddresses
			options.Timeout = time.Second * 5
		},
	)
	return etcdRegistry
}
