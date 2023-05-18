package post

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	pbPost "post/proto"
	"strconv"
)

//  GetCategory getCategory
//	@Summary		获取分区信息 api
//	@Tags			post
//	@Description	获取分区信息
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	internal.Response
//	@Router			/post/category/all [get]
func (p *postHandler) GetCategory(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	resp, err := p.PostService.GetCategory(ctx, &pbPost.Request{})
	if err != nil {
		internal.ServerError(c, errno.GetCategoryInfoError.Error())
		return
	}
	
	internal.Success(c, resp.Categories)
}

type getPostOfCategoryRequest struct {
	CategoryId uint32 `json:"category_id" binding:"required"`
}

//  GetPostCategory getPostCategory
//	@Summary		获取分区帖子信息 api
//	@Tags			post
//	@Description	获取分区帖子信息
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int							false	"页码"
//	@Param			limit	query		int							false	"条数"
//	@Param			object	body		getPostOfCategoryRequest	true	"分类信息"
//	@Success		200		{object}	internal.Response
//	@Router			/post/category [get]
func (p *postHandler) GetPostOfCategory(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	var category getPostOfCategoryRequest
	if err := c.ShouldBindJSON(&category); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	page, errPage := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if errPage != nil || errLimit != nil {
		page = 1
		limit = 10
	}
	
	resp, err := p.PostService.GetPostOfCategory(ctx, &pbPost.PostOfCategoryRequest{
		Category: &pbPost.CategoryInfo{
			CategoryId: category.CategoryId,
		},
		Limit: uint32(limit),
		Page:  uint32(page),
	})
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data)
}
