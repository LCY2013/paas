package tmpl

var (
	WrapperAPI = `package client

import (
	"context"

	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
	{{.Alias}} "path/to/service/proto"
)

type {{.Alias}}Key struct {}

// {{title .Alias}}FromContext FromContext retrieves the client from the Context
func {{title .Alias}}FromContext(ctx context.Context) ({{.Alias}}.{{title .Alias}}Service, bool) {
	c, ok := ctx.Value({{.Alias}}Key{}).({{.Alias}}.{{title .Alias}}Service)
	return c, ok
}

// {{title .Alias}}Wrapper Client returns a wrapper for the {{title .Alias}}Client
func {{title .Alias}}Wrapper(service micro.Service) server.HandlerWrapper {
	client := {{.Alias}}.New{{title .Alias}}Service("go.micro.service.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, {{.Alias}}Key{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
`
)
