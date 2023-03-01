package registry

import (
	"github.com/LCY2013/paas/api-gateway/config"
	"github.com/LCY2013/paas/api-gateway/registry/consul"
	"github.com/LCY2013/paas/api-gateway/registry/etcd"
	"github.com/LCY2013/paas/api-gateway/registry/kubernetes"
	"go-micro.dev/v4/registry"
)

var (
	registryInfo = map[string]func(config *config.Config) registry.Registry{}
)

func Register(config *config.Config) registry.Registry {
	registryFun, ok := registryInfo[config.RegistryType]
	if !ok {
		return registry.DefaultRegistry
	}
	return registryFun(config)
}

func RegisterFunc(registryType string, registryFun func(config *config.Config) registry.Registry) {
	registryInfo[registryType] = registryFun
}

func init() {
	RegisterFunc("consul", consul.NewRegistry)
	RegisterFunc("etcd", etcd.NewRegistry)
	RegisterFunc("kubernetes", kubernetes.NewRegistry)
}
