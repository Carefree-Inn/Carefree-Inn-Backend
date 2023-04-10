// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

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

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	UserRegister(ctx context.Context, in *CCNUInfoRequest, opts ...client.CallOption) (*Response, error)
	UserLogin(ctx context.Context, in *CCNUInfoRequest, opts ...client.CallOption) (*CCNULoginResponse, error)
	GetUserProfile(ctx context.Context, in *Request, opts ...client.CallOption) (*InnUserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, in *InnUserProfileRequest, opts ...client.CallOption) (*Response, error)
	GetBatchUserProfile(ctx context.Context, in *BatchUserProfileRequest, opts ...client.CallOption) (*BatchUserProfileResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) UserRegister(ctx context.Context, in *CCNUInfoRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.UserRegister", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserLogin(ctx context.Context, in *CCNUInfoRequest, opts ...client.CallOption) (*CCNULoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.UserLogin", in)
	out := new(CCNULoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserProfile(ctx context.Context, in *Request, opts ...client.CallOption) (*InnUserProfileResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserProfile", in)
	out := new(InnUserProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserProfile(ctx context.Context, in *InnUserProfileRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserProfile", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetBatchUserProfile(ctx context.Context, in *BatchUserProfileRequest, opts ...client.CallOption) (*BatchUserProfileResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetBatchUserProfile", in)
	out := new(BatchUserProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	UserRegister(context.Context, *CCNUInfoRequest, *Response) error
	UserLogin(context.Context, *CCNUInfoRequest, *CCNULoginResponse) error
	GetUserProfile(context.Context, *Request, *InnUserProfileResponse) error
	UpdateUserProfile(context.Context, *InnUserProfileRequest, *Response) error
	GetBatchUserProfile(context.Context, *BatchUserProfileRequest, *BatchUserProfileResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		UserRegister(ctx context.Context, in *CCNUInfoRequest, out *Response) error
		UserLogin(ctx context.Context, in *CCNUInfoRequest, out *CCNULoginResponse) error
		GetUserProfile(ctx context.Context, in *Request, out *InnUserProfileResponse) error
		UpdateUserProfile(ctx context.Context, in *InnUserProfileRequest, out *Response) error
		GetBatchUserProfile(ctx context.Context, in *BatchUserProfileRequest, out *BatchUserProfileResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) UserRegister(ctx context.Context, in *CCNUInfoRequest, out *Response) error {
	return h.UserHandler.UserRegister(ctx, in, out)
}

func (h *userHandler) UserLogin(ctx context.Context, in *CCNUInfoRequest, out *CCNULoginResponse) error {
	return h.UserHandler.UserLogin(ctx, in, out)
}

func (h *userHandler) GetUserProfile(ctx context.Context, in *Request, out *InnUserProfileResponse) error {
	return h.UserHandler.GetUserProfile(ctx, in, out)
}

func (h *userHandler) UpdateUserProfile(ctx context.Context, in *InnUserProfileRequest, out *Response) error {
	return h.UserHandler.UpdateUserProfile(ctx, in, out)
}

func (h *userHandler) GetBatchUserProfile(ctx context.Context, in *BatchUserProfileRequest, out *BatchUserProfileResponse) error {
	return h.UserHandler.GetBatchUserProfile(ctx, in, out)
}