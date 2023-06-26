package post

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	pb "post/proto"
	"strconv"
)

type createPostRequest struct {
	Category struct {
		CategoryId uint32 `json:"category_id"`
		Title      string `json:"title"`
	} `json:"category" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Tags    []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
}

//  CreatePost createPostRequest
//	@Summary		创建帖子 api
//	@Tags			post
//	@Description	创建帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"用户token"
//	@Param			object			body		createPostRequest	true	"帖子信息"
//	@Success		200				{object}	internal.Response
//	@Router			/post [post]
func (p *postHandler) CreatePost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	
	var tags = make([]*pb.TagInfo, 0, len(req.Tags))
	for _, v := range req.Tags {
		tags = append(tags, &pb.TagInfo{
			Title: v.Tag,
		})
	}
	
	_, err := p.PostService.CreatePost(ctx, &pb.CreatePostRequest{
		Account: account,
		Title:   req.Title,
		Content: req.Content,
		Category: &pb.CategoryInfo{
			Title:      req.Category.Title,
			CategoryId: req.Category.CategoryId,
		},
		Tag: tags,
	})
	if err != nil {
		internal.ServerError(c, errno.CreatePostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.CreatePostError.Error())
		return
	}
	
	internal.Success(c, nil)
}

type deletePostRequest struct {
	PostId uint32 `json:"post_id" binding:"required"`
}

//  DeletePost deletePostRequest
//	@Summary		删除帖子 api
//	@Tags			post
//	@Description	删除帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			post_id			query		int		true	"帖子id"
//	@Success		200				{object}	internal.Response
//	@Router			/post [delete]
func (p *postHandler) DeletePost(c *gin.Context) {
	var post = c.Query("category_id")
	if post == "" {
		internal.Error(c, errno.ParamDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			nil, errno.JsonDataError.Error(),
		)
		return
	}
	
	post_id, err := strconv.Atoi(post)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	
	_, err = p.PostService.DeletePost(ctx, &pb.DeletePostRequest{
		PostId:  uint32(post_id),
		Account: account,
	})
	if err != nil {
		internal.ServerError(c, errno.DeletePostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.DeletePostError.Error())
		return
	}
	
	internal.Success(c, nil)
}

//  GetPostOfUser getPostOfUser
//	@Summary		获取用户发布的帖子 api
//	@Tags			user
//	@Description	获取用户发布的帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
//	@Router			/post/user [get]
func (p *postHandler) GetPostOfUser(c *gin.Context) {
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if errPage != nil || errLimit != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	
	resp, err := p.PostService.GetPostOfUser(ctx, &pb.PostOfUserRequest{
		Account: account,
		Limit:   int32(limit),
		Page:    int32(page),
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts...)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}

//  GetPostOfUserLiked getPostOfUserLiked
//	@Summary		获取用户点赞的帖子 api
//	@Tags			user
//	@Description	获取用户点赞的帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
//	@Router			/post/liked [get]
func (p *postHandler) GetPostOfUserLiked(c *gin.Context) {
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if errPage != nil || errLimit != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	
	resp, err := p.PostService.GetPostOfUserLiked(ctx, &pb.PostOfUserRequest{
		Account: account,
		Limit:   int32(limit),
		Page:    int32(page),
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts...)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}

//  GetPost getPost
//	@Summary		帖子详情 api
//	@Tags			post
//	@Description	帖子详情
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			post_id			query		int		true	"帖子id"
//	@Success		200				{object}	internal.Response
//	@Router			/post/info [get]
func (p *postHandler) GetPost(c *gin.Context) {
	postId := c.DefaultQuery("post_id", "-1")
	if postId == "-1" {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	id, err := strconv.Atoi(postId)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), err)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.GetPost(ctx, &pb.GetPostRequest{PostId: uint32(id)})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Post)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data[0])
}

//  GetPost getPost
//	@Summary		话题广场 api
//	@Tags			tag
//	@Description	话题广场
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Success		200				{object}	internal.Response
//	@Router			/post/square [get]
func (p *postHandler) PostSquare(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.PostSquare(ctx, &pb.Request{})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	tag := make([]string, 0, len(resp.Tags))
	for _, val := range resp.Tags {
		tag = append(tag, val.Title)
	}
	
	internal.Success(c, tag)
}
