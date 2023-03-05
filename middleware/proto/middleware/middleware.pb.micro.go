// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/middleware/middleware.proto

package middleware

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Middleware service

func NewMiddlewareEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Middleware service

type MiddlewareService interface {
	// 对外提供添加服务
	AddMiddleware(ctx context.Context, in *MiddlewareInfo, opts ...client.CallOption) (*Response, error)
	DeleteMiddleware(ctx context.Context, in *MiddlewareId, opts ...client.CallOption) (*Response, error)
	UpdateMiddleware(ctx context.Context, in *MiddlewareInfo, opts ...client.CallOption) (*Response, error)
	FindMiddlewareByID(ctx context.Context, in *MiddlewareId, opts ...client.CallOption) (*MiddlewareInfo, error)
	FindAllMiddleware(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllMiddleware, error)
	// 根据中间件的类型查找所有中间件
	FindAllMiddlewareByTypeID(ctx context.Context, in *FindAllByTypeId, opts ...client.CallOption) (*AllMiddleware, error)
	// 获取中间件类型
	FindMiddleTypeByID(ctx context.Context, in *MiddleTypeId, opts ...client.CallOption) (*MiddleTypeInfo, error)
	AddMiddleType(ctx context.Context, in *MiddleTypeInfo, opts ...client.CallOption) (*Response, error)
	DeleteMiddleTypeByID(ctx context.Context, in *MiddleTypeId, opts ...client.CallOption) (*Response, error)
	UpdateMiddleType(ctx context.Context, in *MiddleTypeInfo, opts ...client.CallOption) (*Response, error)
	FindAllMiddleType(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllMiddleType, error)
}

type middlewareService struct {
	c    client.Client
	name string
}

func NewMiddlewareService(name string, c client.Client) MiddlewareService {
	return &middlewareService{
		c:    c,
		name: name,
	}
}

func (c *middlewareService) AddMiddleware(ctx context.Context, in *MiddlewareInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.AddMiddleware", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) DeleteMiddleware(ctx context.Context, in *MiddlewareId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.DeleteMiddleware", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) UpdateMiddleware(ctx context.Context, in *MiddlewareInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.UpdateMiddleware", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) FindMiddlewareByID(ctx context.Context, in *MiddlewareId, opts ...client.CallOption) (*MiddlewareInfo, error) {
	req := c.c.NewRequest(c.name, "Middleware.FindMiddlewareByID", in)
	out := new(MiddlewareInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) FindAllMiddleware(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllMiddleware, error) {
	req := c.c.NewRequest(c.name, "Middleware.FindAllMiddleware", in)
	out := new(AllMiddleware)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) FindAllMiddlewareByTypeID(ctx context.Context, in *FindAllByTypeId, opts ...client.CallOption) (*AllMiddleware, error) {
	req := c.c.NewRequest(c.name, "Middleware.FindAllMiddlewareByTypeID", in)
	out := new(AllMiddleware)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) FindMiddleTypeByID(ctx context.Context, in *MiddleTypeId, opts ...client.CallOption) (*MiddleTypeInfo, error) {
	req := c.c.NewRequest(c.name, "Middleware.FindMiddleTypeByID", in)
	out := new(MiddleTypeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) AddMiddleType(ctx context.Context, in *MiddleTypeInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.AddMiddleType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) DeleteMiddleTypeByID(ctx context.Context, in *MiddleTypeId, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.DeleteMiddleTypeByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) UpdateMiddleType(ctx context.Context, in *MiddleTypeInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Middleware.UpdateMiddleType", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middlewareService) FindAllMiddleType(ctx context.Context, in *FindAll, opts ...client.CallOption) (*AllMiddleType, error) {
	req := c.c.NewRequest(c.name, "Middleware.FindAllMiddleType", in)
	out := new(AllMiddleType)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Middleware service

type MiddlewareHandler interface {
	// 对外提供添加服务
	AddMiddleware(context.Context, *MiddlewareInfo, *Response) error
	DeleteMiddleware(context.Context, *MiddlewareId, *Response) error
	UpdateMiddleware(context.Context, *MiddlewareInfo, *Response) error
	FindMiddlewareByID(context.Context, *MiddlewareId, *MiddlewareInfo) error
	FindAllMiddleware(context.Context, *FindAll, *AllMiddleware) error
	// 根据中间件的类型查找所有中间件
	FindAllMiddlewareByTypeID(context.Context, *FindAllByTypeId, *AllMiddleware) error
	// 获取中间件类型
	FindMiddleTypeByID(context.Context, *MiddleTypeId, *MiddleTypeInfo) error
	AddMiddleType(context.Context, *MiddleTypeInfo, *Response) error
	DeleteMiddleTypeByID(context.Context, *MiddleTypeId, *Response) error
	UpdateMiddleType(context.Context, *MiddleTypeInfo, *Response) error
	FindAllMiddleType(context.Context, *FindAll, *AllMiddleType) error
}

func RegisterMiddlewareHandler(s server.Server, hdlr MiddlewareHandler, opts ...server.HandlerOption) error {
	type middleware interface {
		AddMiddleware(ctx context.Context, in *MiddlewareInfo, out *Response) error
		DeleteMiddleware(ctx context.Context, in *MiddlewareId, out *Response) error
		UpdateMiddleware(ctx context.Context, in *MiddlewareInfo, out *Response) error
		FindMiddlewareByID(ctx context.Context, in *MiddlewareId, out *MiddlewareInfo) error
		FindAllMiddleware(ctx context.Context, in *FindAll, out *AllMiddleware) error
		FindAllMiddlewareByTypeID(ctx context.Context, in *FindAllByTypeId, out *AllMiddleware) error
		FindMiddleTypeByID(ctx context.Context, in *MiddleTypeId, out *MiddleTypeInfo) error
		AddMiddleType(ctx context.Context, in *MiddleTypeInfo, out *Response) error
		DeleteMiddleTypeByID(ctx context.Context, in *MiddleTypeId, out *Response) error
		UpdateMiddleType(ctx context.Context, in *MiddleTypeInfo, out *Response) error
		FindAllMiddleType(ctx context.Context, in *FindAll, out *AllMiddleType) error
	}
	type Middleware struct {
		middleware
	}
	h := &middlewareHandler{hdlr}
	return s.Handle(s.NewHandler(&Middleware{h}, opts...))
}

type middlewareHandler struct {
	MiddlewareHandler
}

func (h *middlewareHandler) AddMiddleware(ctx context.Context, in *MiddlewareInfo, out *Response) error {
	return h.MiddlewareHandler.AddMiddleware(ctx, in, out)
}

func (h *middlewareHandler) DeleteMiddleware(ctx context.Context, in *MiddlewareId, out *Response) error {
	return h.MiddlewareHandler.DeleteMiddleware(ctx, in, out)
}

func (h *middlewareHandler) UpdateMiddleware(ctx context.Context, in *MiddlewareInfo, out *Response) error {
	return h.MiddlewareHandler.UpdateMiddleware(ctx, in, out)
}

func (h *middlewareHandler) FindMiddlewareByID(ctx context.Context, in *MiddlewareId, out *MiddlewareInfo) error {
	return h.MiddlewareHandler.FindMiddlewareByID(ctx, in, out)
}

func (h *middlewareHandler) FindAllMiddleware(ctx context.Context, in *FindAll, out *AllMiddleware) error {
	return h.MiddlewareHandler.FindAllMiddleware(ctx, in, out)
}

func (h *middlewareHandler) FindAllMiddlewareByTypeID(ctx context.Context, in *FindAllByTypeId, out *AllMiddleware) error {
	return h.MiddlewareHandler.FindAllMiddlewareByTypeID(ctx, in, out)
}

func (h *middlewareHandler) FindMiddleTypeByID(ctx context.Context, in *MiddleTypeId, out *MiddleTypeInfo) error {
	return h.MiddlewareHandler.FindMiddleTypeByID(ctx, in, out)
}

func (h *middlewareHandler) AddMiddleType(ctx context.Context, in *MiddleTypeInfo, out *Response) error {
	return h.MiddlewareHandler.AddMiddleType(ctx, in, out)
}

func (h *middlewareHandler) DeleteMiddleTypeByID(ctx context.Context, in *MiddleTypeId, out *Response) error {
	return h.MiddlewareHandler.DeleteMiddleTypeByID(ctx, in, out)
}

func (h *middlewareHandler) UpdateMiddleType(ctx context.Context, in *MiddleTypeInfo, out *Response) error {
	return h.MiddlewareHandler.UpdateMiddleType(ctx, in, out)
}

func (h *middlewareHandler) FindAllMiddleType(ctx context.Context, in *FindAll, out *AllMiddleType) error {
	return h.MiddlewareHandler.FindAllMiddleType(ctx, in, out)
}
