package registry

import (
	"api/internal/registry/consul"
	"api/internal/registry/etcd"
	"api/internal/registry/kubernetes"
	"go-micro.dev/v4/registry"
)

var (
	registryInfo = map[string]func(op ...registry.Option) registry.Registry{}
)

func Register(registryType string, op ...registry.Option) registry.Registry {
	registryFun, ok := registryInfo[registryType]
	if !ok {
		return registry.DefaultRegistry
	}
	return registryFun(op...)
}

func RegisterFunc(registryType string, registryFun func(op ...registry.Option) registry.Registry) {
	registryInfo[registryType] = registryFun
}

func init() {
	RegisterFunc("consul", consul.NewRegistry)
	RegisterFunc("etcd", etcd.NewRegistry)
	RegisterFunc("kubernetes", kubernetes.NewRegistry)
}
