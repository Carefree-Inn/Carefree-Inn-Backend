// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/post.proto

package post

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

// Api Endpoints for Post service

func NewPostEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Post service

type PostService interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*Response, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...client.CallOption) (*Response, error)
	GetCategory(ctx context.Context, in *Request, opts ...client.CallOption) (*CategoryResponse, error)
	GetPostOfCategory(ctx context.Context, in *PostOfCategoryRequest, opts ...client.CallOption) (*PostResponse, error)
	GetPostOfTag(ctx context.Context, in *PostOfTagRequest, opts ...client.CallOption) (*PostResponse, error)
	GetPostOfUser(ctx context.Context, in *PostOfUserRequest, opts ...client.CallOption) (*PostResponse, error)
	GetPostOfUserLiked(ctx context.Context, in *PostOfUserRequest, opts ...client.CallOption) (*PostResponse, error)
	SearchPost(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*PostResponse, error)
}

type postService struct {
	c    client.Client
	name string
}

func NewPostService(name string, c client.Client) PostService {
	return &postService{
		c:    c,
		name: name,
	}
}

func (c *postService) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Post.CreatePost", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Post.DeletePost", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetCategory(ctx context.Context, in *Request, opts ...client.CallOption) (*CategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetCategory", in)
	out := new(CategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostOfCategory(ctx context.Context, in *PostOfCategoryRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostOfCategory", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostOfTag(ctx context.Context, in *PostOfTagRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostOfTag", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostOfUser(ctx context.Context, in *PostOfUserRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostOfUser", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostOfUserLiked(ctx context.Context, in *PostOfUserRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostOfUserLiked", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) SearchPost(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.SearchPost", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Post service

type PostHandler interface {
	CreatePost(context.Context, *CreatePostRequest, *Response) error
	DeletePost(context.Context, *DeletePostRequest, *Response) error
	GetCategory(context.Context, *Request, *CategoryResponse) error
	GetPostOfCategory(context.Context, *PostOfCategoryRequest, *PostResponse) error
	GetPostOfTag(context.Context, *PostOfTagRequest, *PostResponse) error
	GetPostOfUser(context.Context, *PostOfUserRequest, *PostResponse) error
	GetPostOfUserLiked(context.Context, *PostOfUserRequest, *PostResponse) error
	SearchPost(context.Context, *SearchRequest, *PostResponse) error
}

func RegisterPostHandler(s server.Server, hdlr PostHandler, opts ...server.HandlerOption) error {
	type post interface {
		CreatePost(ctx context.Context, in *CreatePostRequest, out *Response) error
		DeletePost(ctx context.Context, in *DeletePostRequest, out *Response) error
		GetCategory(ctx context.Context, in *Request, out *CategoryResponse) error
		GetPostOfCategory(ctx context.Context, in *PostOfCategoryRequest, out *PostResponse) error
		GetPostOfTag(ctx context.Context, in *PostOfTagRequest, out *PostResponse) error
		GetPostOfUser(ctx context.Context, in *PostOfUserRequest, out *PostResponse) error
		GetPostOfUserLiked(ctx context.Context, in *PostOfUserRequest, out *PostResponse) error
		SearchPost(ctx context.Context, in *SearchRequest, out *PostResponse) error
	}
	type Post struct {
		post
	}
	h := &postHandler{hdlr}
	return s.Handle(s.NewHandler(&Post{h}, opts...))
}

type postHandler struct {
	PostHandler
}

func (h *postHandler) CreatePost(ctx context.Context, in *CreatePostRequest, out *Response) error {
	return h.PostHandler.CreatePost(ctx, in, out)
}

func (h *postHandler) DeletePost(ctx context.Context, in *DeletePostRequest, out *Response) error {
	return h.PostHandler.DeletePost(ctx, in, out)
}

func (h *postHandler) GetCategory(ctx context.Context, in *Request, out *CategoryResponse) error {
	return h.PostHandler.GetCategory(ctx, in, out)
}

func (h *postHandler) GetPostOfCategory(ctx context.Context, in *PostOfCategoryRequest, out *PostResponse) error {
	return h.PostHandler.GetPostOfCategory(ctx, in, out)
}

func (h *postHandler) GetPostOfTag(ctx context.Context, in *PostOfTagRequest, out *PostResponse) error {
	return h.PostHandler.GetPostOfTag(ctx, in, out)
}

func (h *postHandler) GetPostOfUser(ctx context.Context, in *PostOfUserRequest, out *PostResponse) error {
	return h.PostHandler.GetPostOfUser(ctx, in, out)
}

func (h *postHandler) GetPostOfUserLiked(ctx context.Context, in *PostOfUserRequest, out *PostResponse) error {
	return h.PostHandler.GetPostOfUserLiked(ctx, in, out)
}

func (h *postHandler) SearchPost(ctx context.Context, in *SearchRequest, out *PostResponse) error {
	return h.PostHandler.SearchPost(ctx, in, out)
}
