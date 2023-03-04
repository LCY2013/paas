package consul

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func NewRegistry(op ...registry.Option) registry.Registry {
	var consulRegistry = consul.NewRegistry(op...)

	return consulRegistry
}
