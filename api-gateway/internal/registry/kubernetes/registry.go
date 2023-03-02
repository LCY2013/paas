package kubernetes

import (
	"github.com/go-micro/plugins/v4/registry/kubernetes"
	"go-micro.dev/v4/registry"
)

func NewRegistry(op ...registry.Option) registry.Registry {
	var etcdRegistry = kubernetes.NewRegistry(op...)
	return etcdRegistry
}
