package etcd

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

func NewRegistry(op ...registry.Option) registry.Registry {
	var etcdRegistry = etcd.NewRegistry(op...)
	return etcdRegistry
}
