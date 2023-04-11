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
)

type createPost struct {
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

//  CreatePost createPost
//	@Summary		创建帖子 api
//	@Tags			post
//	@Description	创建帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//  @Param           object     body   createPost true "帖子信息"
//	@Success		200		{object}	internal.Response
//	@Router			/post [post]
func (p *postHandler) CreatePost(c *gin.Context) {
	var req createPost
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

type deletePost struct {
	PostId uint32 `json:"post_id"`
}

//  DeletePost deletePost
//	@Summary		删除帖子 api
//	@Tags			post
//	@Description	删除帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//  @Param           object     body   deletePost true "帖子信息"
//	@Success		200		{object}	internal.Response
//	@Router			/post [delete]
func (p *postHandler) DeletePost(c *gin.Context) {
	var req deletePost
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
