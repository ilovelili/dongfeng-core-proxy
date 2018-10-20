// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng/core-proxy/services/proto/api.proto

/*
Package dongfeng_svc_core_proxy is a generated protocol buffer package.

It is generated from these files:
	github.com/ilovelili/dongfeng/core-proxy/services/proto/api.proto

It has these top-level messages:
	HealthcheckRequest
	HealthcheckResponse
	LoginRequest
	LoginResponse
	UserProfile
	Setting
	Operation
*/
package dongfeng_svc_core_proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Api service

type ApiService interface {
	Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...client.CallOption) (*HealthcheckResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "dongfeng.svc.core.proxy"
	}
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) Healthcheck(ctx context.Context, in *HealthcheckRequest, opts ...client.CallOption) (*HealthcheckResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Healthcheck", in)
	out := new(HealthcheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	Healthcheck(context.Context, *HealthcheckRequest, *HealthcheckResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) {
	type api interface {
		Healthcheck(ctx context.Context, in *HealthcheckRequest, out *HealthcheckResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) Healthcheck(ctx context.Context, in *HealthcheckRequest, out *HealthcheckResponse) error {
	return h.ApiHandler.Healthcheck(ctx, in, out)
}

func (h *apiHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.ApiHandler.Login(ctx, in, out)
}
