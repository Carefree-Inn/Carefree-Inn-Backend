package post

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	pb "github.com/jackj-ohn1/package/proto/post"
	"github.com/pkg/errors"
)

func (p *postHandler) GetPostOfTag(c *gin.Context) {
	var tag tagInfo
	if err := c.ShouldBindJSON(&tag); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.GetPostOfTag(ctx, &pb.TagInfo{
		Title: tag.Title,
		TagId: tag.TagId,
	})
	if err != nil {
		internal.ServerError(c, errno.CreatePostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.CreatePostError.Error())
		return
	}
	
	var accounts = make([]string, 0, len(resp.Posts))
	for _, v := range resp.Posts {
		accounts = append(accounts, v.Account)
	}
	
	data, err := p.GetUserInfoWithAny(ctx, accounts, resp.Posts, p.assemble)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data.([]*PostInfo))
}
