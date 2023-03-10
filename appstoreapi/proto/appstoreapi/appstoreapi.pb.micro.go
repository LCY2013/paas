// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/appstoreapi/appstoreapi.proto

package appstoreapi

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

// Api Endpoints for Appstoreapi service

func NewAppstoreapiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Appstoreapi service

type AppstoreapiService interface {
	FindAppstoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddAppstore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteAppstoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateAppstore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 默认接口
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 添加常用的统计接口
	AddInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type appstoreapiService struct {
	c    client.Client
	name string
}

func NewAppstoreapiService(name string, c client.Client) AppstoreapiService {
	return &appstoreapiService{
		c:    c,
		name: name,
	}
}

func (c *appstoreapiService) FindAppstoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.FindAppstoreById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) AddAppstore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.AddAppstore", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) DeleteAppstoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.DeleteAppstoreById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) UpdateAppstore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.UpdateAppstore", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) AddInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.AddInstallNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) GetInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.GetInstallNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) AddViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.AddViewNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appstoreapiService) GetViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appstoreapi.GetViewNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Appstoreapi service

type AppstoreapiHandler interface {
	FindAppstoreById(context.Context, *Request, *Response) error
	AddAppstore(context.Context, *Request, *Response) error
	DeleteAppstoreById(context.Context, *Request, *Response) error
	UpdateAppstore(context.Context, *Request, *Response) error
	// 默认接口
	Call(context.Context, *Request, *Response) error
	// 添加常用的统计接口
	AddInstallNum(context.Context, *Request, *Response) error
	GetInstallNum(context.Context, *Request, *Response) error
	AddViewNum(context.Context, *Request, *Response) error
	GetViewNum(context.Context, *Request, *Response) error
}

func RegisterAppstoreapiHandler(s server.Server, hdlr AppstoreapiHandler, opts ...server.HandlerOption) error {
	type appstoreapi interface {
		FindAppstoreById(ctx context.Context, in *Request, out *Response) error
		AddAppstore(ctx context.Context, in *Request, out *Response) error
		DeleteAppstoreById(ctx context.Context, in *Request, out *Response) error
		UpdateAppstore(ctx context.Context, in *Request, out *Response) error
		Call(ctx context.Context, in *Request, out *Response) error
		AddInstallNum(ctx context.Context, in *Request, out *Response) error
		GetInstallNum(ctx context.Context, in *Request, out *Response) error
		AddViewNum(ctx context.Context, in *Request, out *Response) error
		GetViewNum(ctx context.Context, in *Request, out *Response) error
	}
	type Appstoreapi struct {
		appstoreapi
	}
	h := &appstoreapiHandler{hdlr}
	return s.Handle(s.NewHandler(&Appstoreapi{h}, opts...))
}

type appstoreapiHandler struct {
	AppstoreapiHandler
}

func (h *appstoreapiHandler) FindAppstoreById(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.FindAppstoreById(ctx, in, out)
}

func (h *appstoreapiHandler) AddAppstore(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.AddAppstore(ctx, in, out)
}

func (h *appstoreapiHandler) DeleteAppstoreById(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.DeleteAppstoreById(ctx, in, out)
}

func (h *appstoreapiHandler) UpdateAppstore(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.UpdateAppstore(ctx, in, out)
}

func (h *appstoreapiHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.Call(ctx, in, out)
}

func (h *appstoreapiHandler) AddInstallNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.AddInstallNum(ctx, in, out)
}

func (h *appstoreapiHandler) GetInstallNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.GetInstallNum(ctx, in, out)
}

func (h *appstoreapiHandler) AddViewNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.AddViewNum(ctx, in, out)
}

func (h *appstoreapiHandler) GetViewNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppstoreapiHandler.GetViewNum(ctx, in, out)
}
