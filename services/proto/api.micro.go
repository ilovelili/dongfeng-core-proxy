// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/ilovelili/dongfeng-core-proxy/services/proto/api.proto

/*
Package dongfeng_svc_core_proxy is a generated protocol buffer package.

It is generated from these files:
	github.com/ilovelili/dongfeng-core-proxy/services/proto/api.proto

It has these top-level messages:
*/
package dongfeng_svc_core_proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dongfeng_protobuf "github.com/ilovelili/dongfeng-protobuf"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = dongfeng_protobuf.GetMenuResponse{}

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
	Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, opts ...client.CallOption) (*dongfeng_protobuf.LoginResponse, error)
	Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, opts ...client.CallOption) (*dongfeng_protobuf.DashboardResponse, error)
	UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateUserResponse, error)
	GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetAttendanceResponse, error)
	UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateAttendanceResponse, error)
	GetNamelist(ctx context.Context, in *dongfeng_protobuf.GetNamelistRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetNamelistResponse, error)
	UpdatePhysique(ctx context.Context, in *dongfeng_protobuf.UpdatePhysiqueRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdatePhysiqueResponse, error)
	GetRecipe(ctx context.Context, in *dongfeng_protobuf.GetRecipeRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *dongfeng_protobuf.UpdateRecipeRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateRecipeResponse, error)
	GetIngredient(ctx context.Context, in *dongfeng_protobuf.GetIngredientRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetIngredientResponse, error)
	UpdateIngredient(ctx context.Context, in *dongfeng_protobuf.UpdateIngredientRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateIngredientResponse, error)
	GetMenu(ctx context.Context, in *dongfeng_protobuf.GetMenuRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetMenuResponse, error)
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

func (c *apiService) Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, opts ...client.CallOption) (*dongfeng_protobuf.LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Login", in)
	out := new(dongfeng_protobuf.LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, opts ...client.CallOption) (*dongfeng_protobuf.DashboardResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Dashboard", in)
	out := new(dongfeng_protobuf.DashboardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateUserResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateUser", in)
	out := new(dongfeng_protobuf.UpdateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetAttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetAttendance", in)
	out := new(dongfeng_protobuf.GetAttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateAttendanceResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateAttendance", in)
	out := new(dongfeng_protobuf.UpdateAttendanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetNamelist(ctx context.Context, in *dongfeng_protobuf.GetNamelistRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetNamelistResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetNamelist", in)
	out := new(dongfeng_protobuf.GetNamelistResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdatePhysique(ctx context.Context, in *dongfeng_protobuf.UpdatePhysiqueRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdatePhysiqueResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdatePhysique", in)
	out := new(dongfeng_protobuf.UpdatePhysiqueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetRecipe(ctx context.Context, in *dongfeng_protobuf.GetRecipeRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetRecipeResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetRecipe", in)
	out := new(dongfeng_protobuf.GetRecipeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateRecipe(ctx context.Context, in *dongfeng_protobuf.UpdateRecipeRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateRecipeResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateRecipe", in)
	out := new(dongfeng_protobuf.UpdateRecipeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetIngredient(ctx context.Context, in *dongfeng_protobuf.GetIngredientRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetIngredientResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetIngredient", in)
	out := new(dongfeng_protobuf.GetIngredientResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) UpdateIngredient(ctx context.Context, in *dongfeng_protobuf.UpdateIngredientRequest, opts ...client.CallOption) (*dongfeng_protobuf.UpdateIngredientResponse, error) {
	req := c.c.NewRequest(c.name, "Api.UpdateIngredient", in)
	out := new(dongfeng_protobuf.UpdateIngredientResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiService) GetMenu(ctx context.Context, in *dongfeng_protobuf.GetMenuRequest, opts ...client.CallOption) (*dongfeng_protobuf.GetMenuResponse, error) {
	req := c.c.NewRequest(c.name, "Api.GetMenu", in)
	out := new(dongfeng_protobuf.GetMenuResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	Login(context.Context, *dongfeng_protobuf.LoginRequest, *dongfeng_protobuf.LoginResponse) error
	Dashboard(context.Context, *dongfeng_protobuf.DashboardRequest, *dongfeng_protobuf.DashboardResponse) error
	UpdateUser(context.Context, *dongfeng_protobuf.UpdateUserRequest, *dongfeng_protobuf.UpdateUserResponse) error
	GetAttendance(context.Context, *dongfeng_protobuf.GetAttendanceRequest, *dongfeng_protobuf.GetAttendanceResponse) error
	UpdateAttendance(context.Context, *dongfeng_protobuf.UpdateAttendanceRequest, *dongfeng_protobuf.UpdateAttendanceResponse) error
	GetNamelist(context.Context, *dongfeng_protobuf.GetNamelistRequest, *dongfeng_protobuf.GetNamelistResponse) error
	UpdatePhysique(context.Context, *dongfeng_protobuf.UpdatePhysiqueRequest, *dongfeng_protobuf.UpdatePhysiqueResponse) error
	GetRecipe(context.Context, *dongfeng_protobuf.GetRecipeRequest, *dongfeng_protobuf.GetRecipeResponse) error
	UpdateRecipe(context.Context, *dongfeng_protobuf.UpdateRecipeRequest, *dongfeng_protobuf.UpdateRecipeResponse) error
	GetIngredient(context.Context, *dongfeng_protobuf.GetIngredientRequest, *dongfeng_protobuf.GetIngredientResponse) error
	UpdateIngredient(context.Context, *dongfeng_protobuf.UpdateIngredientRequest, *dongfeng_protobuf.UpdateIngredientResponse) error
	GetMenu(context.Context, *dongfeng_protobuf.GetMenuRequest, *dongfeng_protobuf.GetMenuResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) {
	type api interface {
		Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, out *dongfeng_protobuf.LoginResponse) error
		Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, out *dongfeng_protobuf.DashboardResponse) error
		UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, out *dongfeng_protobuf.UpdateUserResponse) error
		GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, out *dongfeng_protobuf.GetAttendanceResponse) error
		UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, out *dongfeng_protobuf.UpdateAttendanceResponse) error
		GetNamelist(ctx context.Context, in *dongfeng_protobuf.GetNamelistRequest, out *dongfeng_protobuf.GetNamelistResponse) error
		UpdatePhysique(ctx context.Context, in *dongfeng_protobuf.UpdatePhysiqueRequest, out *dongfeng_protobuf.UpdatePhysiqueResponse) error
		GetRecipe(ctx context.Context, in *dongfeng_protobuf.GetRecipeRequest, out *dongfeng_protobuf.GetRecipeResponse) error
		UpdateRecipe(ctx context.Context, in *dongfeng_protobuf.UpdateRecipeRequest, out *dongfeng_protobuf.UpdateRecipeResponse) error
		GetIngredient(ctx context.Context, in *dongfeng_protobuf.GetIngredientRequest, out *dongfeng_protobuf.GetIngredientResponse) error
		UpdateIngredient(ctx context.Context, in *dongfeng_protobuf.UpdateIngredientRequest, out *dongfeng_protobuf.UpdateIngredientResponse) error
		GetMenu(ctx context.Context, in *dongfeng_protobuf.GetMenuRequest, out *dongfeng_protobuf.GetMenuResponse) error
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

func (h *apiHandler) Login(ctx context.Context, in *dongfeng_protobuf.LoginRequest, out *dongfeng_protobuf.LoginResponse) error {
	return h.ApiHandler.Login(ctx, in, out)
}

func (h *apiHandler) Dashboard(ctx context.Context, in *dongfeng_protobuf.DashboardRequest, out *dongfeng_protobuf.DashboardResponse) error {
	return h.ApiHandler.Dashboard(ctx, in, out)
}

func (h *apiHandler) UpdateUser(ctx context.Context, in *dongfeng_protobuf.UpdateUserRequest, out *dongfeng_protobuf.UpdateUserResponse) error {
	return h.ApiHandler.UpdateUser(ctx, in, out)
}

func (h *apiHandler) GetAttendance(ctx context.Context, in *dongfeng_protobuf.GetAttendanceRequest, out *dongfeng_protobuf.GetAttendanceResponse) error {
	return h.ApiHandler.GetAttendance(ctx, in, out)
}

func (h *apiHandler) UpdateAttendance(ctx context.Context, in *dongfeng_protobuf.UpdateAttendanceRequest, out *dongfeng_protobuf.UpdateAttendanceResponse) error {
	return h.ApiHandler.UpdateAttendance(ctx, in, out)
}

func (h *apiHandler) GetNamelist(ctx context.Context, in *dongfeng_protobuf.GetNamelistRequest, out *dongfeng_protobuf.GetNamelistResponse) error {
	return h.ApiHandler.GetNamelist(ctx, in, out)
}

func (h *apiHandler) UpdatePhysique(ctx context.Context, in *dongfeng_protobuf.UpdatePhysiqueRequest, out *dongfeng_protobuf.UpdatePhysiqueResponse) error {
	return h.ApiHandler.UpdatePhysique(ctx, in, out)
}

func (h *apiHandler) GetRecipe(ctx context.Context, in *dongfeng_protobuf.GetRecipeRequest, out *dongfeng_protobuf.GetRecipeResponse) error {
	return h.ApiHandler.GetRecipe(ctx, in, out)
}

func (h *apiHandler) UpdateRecipe(ctx context.Context, in *dongfeng_protobuf.UpdateRecipeRequest, out *dongfeng_protobuf.UpdateRecipeResponse) error {
	return h.ApiHandler.UpdateRecipe(ctx, in, out)
}

func (h *apiHandler) GetIngredient(ctx context.Context, in *dongfeng_protobuf.GetIngredientRequest, out *dongfeng_protobuf.GetIngredientResponse) error {
	return h.ApiHandler.GetIngredient(ctx, in, out)
}

func (h *apiHandler) UpdateIngredient(ctx context.Context, in *dongfeng_protobuf.UpdateIngredientRequest, out *dongfeng_protobuf.UpdateIngredientResponse) error {
	return h.ApiHandler.UpdateIngredient(ctx, in, out)
}

func (h *apiHandler) GetMenu(ctx context.Context, in *dongfeng_protobuf.GetMenuRequest, out *dongfeng_protobuf.GetMenuResponse) error {
	return h.ApiHandler.GetMenu(ctx, in, out)
}
