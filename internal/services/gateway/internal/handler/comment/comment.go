package comment

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"strconv"
	"user-post/pkg/log"
	pb "user-post/proto"
	pbUser "user/proto"
)

type makeCommentRequest struct {
	IsTop        bool   `json:"is_top"`
	TopCommentId uint32 `json:"top_comment_id"`
	
	ToUserAccount string `json:"to_user_account"`
	Content       string `json:"content"`
	PostId        uint32 `json:"post_id"`
	
	FromUserNickName string `json:"from_user_nickname"`
	FromUserAvatar   string `json:"from_user_avatar"`
}

//  MakeComment makeComment
//	@Summary		评论 api
//	@Tags			comment
//	@Description	评论
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"用户token"
//	@Param			object			body		makeCommentRequest	true	"评论信息"
//	@Success		200				{object}	internal.Response
//	@Router			/comment [post]
func (l *commentHandler) MakeComment(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	var req makeCommentRequest
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	if req.IsTop && req.ToUserAccount != "" ||
		!req.IsTop && req.ToUserAccount == "" {
		internal.Error(c, errno.ConstraintParamError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ConstraintParamError)
		return
	}
	
	account := c.MustGet("account").(string)
	
	_, err := l.UserPostService.MakeComment(ctx, &pb.MakeCommentRequest{
		PostId:           req.PostId,
		IsTop:            req.IsTop,
		ToUserAccount:    req.ToUserAccount,
		FromUserAccount:  account,
		Content:          req.Content,
		TopCommentId:     req.TopCommentId,
		FromUserAvatar:   req.FromUserAvatar,
		FromUserNickName: req.FromUserNickName,
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, "评论成功")
}

//type deleteCommentRequest struct {
//	CommentId uint32 `json:"comment_id"`
//}

//  DeleteComment deleteComment
//	@Summary		删除评论 api
//	@Tags			comment
//	@Description	删除评论
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			comment_id		query		int		true	"评论id"
//	@Success		200				{object}	internal.Response
//	@Router			/comment [delete]
func (l *commentHandler) DeleteComment(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	
	var comment = c.Query("comment_id")
	if comment == "" {
		internal.Error(c, errno.ParamDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			nil, errno.JsonDataError.Error(),
		)
		return
	}
	
	commentId, err := strconv.Atoi(comment)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	_, err = l.UserPostService.DeleteComment(ctx, &pb.DeleteCommentRequest{
		CommentId: uint32(commentId),
	})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, "删除评论成功")
}

//  GetCommentOfPost getCommentOfPost
//	@Summary		获取帖子下的评论 api
//	@Tags			comment
//	@Description	获取帖子下的评论
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			post_id			query		int		true	"帖子id"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
//	@Router			/comment/post [get]
func (l *commentHandler) GetCommentOfPost(c *gin.Context) {
	pageStr, limitStr := c.DefaultQuery("page", "1"), c.DefaultQuery("limit", "10")
	postIdStr := c.Query("post_id")
	
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
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
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := l.UserPostService.GetCommentOfPost(ctx,
		&pb.GetCommentOfPostRequest{
			PostId: uint32(postId),
			Limit:  uint32(limit),
			Page:   uint32(page),
		})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	data, err := l.UpperAssemble(ctx, resp)
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}

//  GetCommentOfUser getCommentOfUser
//	@Summary		获取用户的评论 api
//	@Tags			user
//	@Description	获取用户的评论
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
//	@Router			/comment/user [get]
func (l *commentHandler) GetCommentOfUser(c *gin.Context) {
	pageStr, limitStr := c.DefaultQuery("page", "1"), c.DefaultQuery("limit", "10")
	
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	resp, err := l.UserPostService.GetCommentOfUser(ctx,
		&pb.GetCommentOfUserRequest{
			Account: account,
			Limit:   int32(uint32(limit)),
			Page:    int32(uint32(page)),
		})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	selfInfo, err := l.UserService.GetUserProfile(ctx, &pbUser.Request{Account: account})
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	data := make([]*Comment, 0, len(resp.Comments))
	for _, val := range resp.Comments {
		one := &Comment{
			CommentId: val.CommentId,
			FromUserAccount: &UserInfo{
				Account:  selfInfo.Account,
				Avatar:   selfInfo.Avatar,
				Sex:      int8(selfInfo.Sex),
				Nickname: selfInfo.Nickname,
			},
			ToUserAccount: &UserInfo{Account: val.ToUserId},
			Content:       val.Content,
			CreateTime:    val.CreateTime,
			
			PostId:       val.PostId,
			IsTop:        val.IsTop,
			TopCommentId: val.TopCommentId,
		}
		data = append(data, one)
	}
	
	internal.Success(c, data)
}
