package post

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	pb "post/proto"
)

type searchRequest struct {
	SearchType string `json:"search_type" binding:"required"`
	Data       string `json:"data" binding:"required"`
}

//  SearchPost searchPost
//	@Summary		搜索帖子 api
//	@Tags			post
//	@Description	搜索帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			true	"用户token"
//	@Param			object			body		searchRequest	true	"搜索信息"
//	@Success		200				{object}	internal.Response
//	@Router			/post/search [post]
func (p *postHandler) SearchPost(c *gin.Context) {
	var req searchRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	account := c.MustGet("account").(string)
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.SearchPost(ctx, &pb.SearchRequest{
		SearchType: req.SearchType,
		Content:    req.Data,
		Account:    account,
	})
	if err != nil {
		if errno.Is(err, errno.ResourceNotExist) {
			log.Warn(
				log.WithField("X-Request-Id", c.MustGet("uuid")),
				errors.WithStack(err), err.Error(),
			)
			internal.Error(c, err)
			return
		}
		internal.ServerError(c, errno.DatabaseError.Error())
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts...)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data)
}
