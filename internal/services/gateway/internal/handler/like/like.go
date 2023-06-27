package like

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
	pb "user-post/proto"
)

type makeLikeRequest struct {
	FromUserAvatar   string `json:"from_user_avatar"`
	FromUserNickname string `json:"from_user_nickname"`
	
	ToUserAccount string `json:"to_user_account"`
	PostId        int    `json:"post_id"`
}

//  MakeLike makeLike
//	@Summary		点赞 api
//	@Tags			like
//	@Description	点赞
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			true	"用户token"
//	@Param			object			body		makeLikeRequest	true	"被点赞帖子相关信息"
//	@Success		200				{object}	internal.Response
//	@Router			/like [post]
func (l *likeHandler) MakeLike(c *gin.Context) {
	var req makeLikeRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	_, err := l.UserPostService.MakeLike(ctx, &pb.MakeLikeRequest{
		PostId:           uint32(req.PostId),
		FromUserAccount:  c.MustGet("account").(string),
		FromUserAvatar:   req.FromUserAvatar,
		FromUserNickname: req.FromUserNickname,
		ToUserAccount:    req.ToUserAccount,
	})
	if err != nil {
		if errno.Is(err, errno.DuplicateLike) {
			internal.Error(c, errno.DuplicateLike)
			return
		}
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, "点赞成功")
}

//  CancelLike cancelLike
//	@Summary		取消点赞 api
//	@Tags			like
//	@Description	取消点赞
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			post_id			query		int		true	"帖子id"
//	@Success		200				{object}	internal.Response
//	@Router			/like [delete]
func (l *likeHandler) CancelLike(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	postIdStr := c.Query("post_id")
	if postIdStr == "" {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	_, err = l.UserPostService.CancelLike(ctx, &pb.CancelLikeRequest{
		PostId:  uint32(postId),
		Account: c.MustGet("account").(string),
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, "取消点赞成功")
}

//func (l *likeHandler) GetLikeOfUser(c *gin.Context) {
//	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
//
//	account := c.MustGet("account").(string)
//	pageStr, limitStr := c.DefaultQuery("page", "1"), c.DefaultQuery("limit", "10")
//
//	page, err := strconv.Atoi(pageStr)
//	if err != nil {
//		internal.Error(c, errno.ParamDataError)
//		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
//		return
//	}
//	limit, err := strconv.Atoi(limitStr)
//	if err != nil {
//		internal.Error(c, errno.ParamDataError)
//		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
//		return
//	}
//
//	_, err = l.UserPostService.GetLikeOfUser(ctx, &pb.LikeOfUserRequest{
//		Account: account,
//		Page:    int32(page),
//		Limit:   int32(limit),
//	})
//	if err != nil {
//		internal.ServerError(c, errno.InternalServerError.Error())
//		return
//	}
//
//	internal.Success(c, "")
//}
//
//type isBatchLikedRequest struct {
//	PostIds []uint32 `json:"post_ids"`
//}
//
//func (l *likeHandler) IsBatchLiked(c *gin.Context) {
//	var req isBatchLikedRequest
//	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
//		internal.Error(c, errno.JsonDataError)
//		log.Warn(
//			log.WithField("X-Request-Id", c.MustGet("uuid")),
//			errors.WithStack(err), errno.JsonDataError.Error(),
//		)
//		return
//	}
//
//	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
//	account := c.MustGet("account").(string)
//
//	resp, err := l.UserPostService.IsBatchLiked(ctx, &pb.BactchLikedRequest{
//		Account: account,
//		PostIds: req.PostIds,
//	})
//	if err != nil {
//		internal.ServerError(c, errno.InternalServerError.Error())
//		return
//	}
//
//	internal.Success(c, resp.Status)
//}
