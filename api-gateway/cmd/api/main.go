package main

import (
	"context"
	"github.com/LCY2013/paas/api-gateway/config"
	agr "github.com/LCY2013/paas/api-gateway/registry"
	"github.com/LCY2013/paas/common/logger"
	"go-micro.dev/v4/api"
	"go-micro.dev/v4/api/router"
	"go-micro.dev/v4/api/router/registry"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := api.NewApi(
		api.WithRouter(registry.NewRouter(router.WithRegistry(agr.Register(config.Conf())))),
	)

	if err := srv.Run(ctx); err != nil {
		logger.Fatal(err)
	}
}
