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
//	@Param			Authorzation	header		string				true	"用户token"
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
//	@Param			Authorzation	header		string				true	"用户token"
//	@Param			object			body		deletePostRequest	true	"帖子信息"
//	@Success		200				{object}	internal.Response
//	@Router			/post [delete]
func (p *postHandler) DeletePost(c *gin.Context) {
	var req deletePostRequest
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
	
	_, err := p.PostService.DeletePost(ctx, &pb.DeletePostRequest{
		PostId:  req.PostId,
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
//	@Tags			post
//	@Description	获取用户发布的帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
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
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}

//  GetPostOfUserLiked getPostOfUserLiked
//	@Summary		获取用户点赞的帖子 api
//	@Tags			post
//	@Description	获取用户点赞的帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
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
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}
