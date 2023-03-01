package consul

import (
	"github.com/LCY2013/paas/api-gateway/config"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"time"
)

func NewRegistry(config *config.Config) registry.Registry {
	var consulRegistry = consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = config.RegisterAddresses
		options.Timeout = time.Second * 5
	})

	return consulRegistry
}
