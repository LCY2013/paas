package kubernetes

import (
	"github.com/LCY2013/paas/api-gateway/config"
	"github.com/go-micro/plugins/v4/registry/kubernetes"
	"go-micro.dev/v4/registry"
	"time"
)

func NewRegistry(config *config.Config) registry.Registry {
	var etcdRegistry = kubernetes.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = config.RegisterAddresses
			options.Timeout = time.Second * 5
		},
	)
	return etcdRegistry
}
