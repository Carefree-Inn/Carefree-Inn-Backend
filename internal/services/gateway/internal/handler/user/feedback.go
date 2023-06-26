package user

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	pb "user/proto"
)

type makeFeedbackRequest struct {
	PostId       int32  `json:"post_id"`
	FeedbackType string `json:"feedback_type"`
	Content      string `json:"content"`
}

//  MakeFeedback MakeFeedback
//	@Summary		反馈 api
//	@Tags			user
//	@Description	用户对帖子的反馈信息
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"用户token"
//	@Param			object			body		makeFeedbackRequest	true	"反馈信息"
//	@Success		200				{object}	internal.Response
//	@Router			/user/feedback [post]
func (u *userHandler) MakeFeedback(c *gin.Context) {
	var req = makeFeedbackRequest{}
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
	
	_, err := u.UserService.MakeFeedback(ctx, &pb.MakeFeedbackRequest{
		Account:      account,
		Content:      req.Content,
		PostId:       req.PostId,
		FeedbackType: req.FeedbackType,
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		log.Warn(nil, errors.WithStack(err), err.Error())
		return
	}
	
	internal.Success(c, nil)
}
