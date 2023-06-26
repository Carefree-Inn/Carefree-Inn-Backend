package post

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	pb "post/proto"
)

//type getPostOfTag struct {
//	Title string `json:"title" binding:"required"`
//}

//  GetPostOfTag getPostOfTag
//	@Summary		获取tag下的帖子 api
//	@Tags			tag
//	@Description	获取tag下的帖子
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			tag				query		int		true	"tag标题"
//	@Success		200				{object}	internal.Response
//	@Router			/post/tag [get]
func (p *postHandler) GetPostOfTag(c *gin.Context) {
	var tag = c.Query("tag")
	if tag == "" {
		internal.Error(c, errno.ParamDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			nil, errno.JsonDataError.Error(),
		)
		return
	}
	
	account, exist := c.Get("account")
	if !exist {
		account = ""
	}
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := p.PostService.GetPostOfTag(ctx, &pb.PostOfTagRequest{
		Title:   tag,
		Account: account.(string),
	})
	if err != nil {
		internal.ServerError(c, errno.CreatePostError.Error())
		return
	}
	
	var accounts = make([]string, 0, len(resp.Posts))
	for _, v := range resp.Posts {
		accounts = append(accounts, v.Account)
	}
	
	data, err := p.AssemblePostAndUser(ctx, resp.Posts...)
	if err != nil {
		internal.ServerError(c, errno.GetCategoryCategoryPostError.Error())
		return
	}
	
	internal.Success(c, data)
}
