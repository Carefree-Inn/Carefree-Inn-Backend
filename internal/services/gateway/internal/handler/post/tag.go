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

type getPostOfTag struct {
	Title string `json:"title" binding:"required"`
}

//  GetPostOfTag getPostOfTag
//	@Summary		获取tag下的帖子 api
//	@Tags			post
//	@Description	获取tag下的帖子
//	@Accept			json
//	@Produce		json
//	@Param			object	body		getPostOfTag	true	"tag信息"
//	@Success		200		{object}	internal.Response
//	@Router			/post/tag [post]
func (p *postHandler) GetPostOfTag(c *gin.Context) {
	var tag getPostOfTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.GetPostOfTag(ctx, &pb.PostOfTagRequest{
		Title: tag.Title,
	})
	if err != nil {
		internal.ServerError(c, errno.CreatePostError.Error())
		return
	}
	
	var accounts = make([]string, 0, len(resp.Posts))
	for _, v := range resp.Posts {
		accounts = append(accounts, v.Account)
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data)
}
