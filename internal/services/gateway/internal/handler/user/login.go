package user

import (
	"context"
	"gateway/internal"
	"gateway/pkg"
	errno "gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	pb "user/proto"
)

type user struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Sex      int8   `json:"sex"`
}

//  Register register
//	@Summary		注册 api
//	@Tags			user
//	@Description	用户通过学号注册
//	@Accept			json
//	@Produce		json
//	@Param			object	body		user	true	"用户信息"
//	@Success		200		{object}	internal.Response
//	@Router			/user/register [post]
func (u *userHandler) Register(c *gin.Context) {
	var req user
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	_, err := u.UserRegister(ctx, &pb.CCNUInfoRequest{
		Account:  req.Account,
		Password: req.Password,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Sex:      pb.Sex(req.Sex),
	})
	if err != nil {
		if errno.Is(err, errno.UserNotExistError) {
			internal.Error(c, err)
			return
		} else if errno.Is(err, errno.LoginWrongInfoError) {
			internal.Error(c, err)
			return
		}
		
		internal.ServerError(c, errno.LoginServerError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			err, err.Error())
		return
	}
	
	str, err := pkg.GenerateToken(req.Account)
	if err != nil {
		internal.ServerError(c, errno.TokenGenerateError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err))
		return
	}
	
	internal.Success(c, str)
	
}

//  Login login
//	@Summary		登陆 api
//	@Tags			user
//	@Description	用户通过学号登录
//	@Accept			json
//	@Produce		json
//	@Param			object	body		user	true	"用户信息"
//	@Success		200		{object}	internal.Response
//	@Router			/user/login [post]
func (u *userHandler) Login(c *gin.Context) {
	var req user
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Warn(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	_, err := u.UserLogin(ctx, &pb.CCNUInfoRequest{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		if errno.Is(err, errno.UserNotExistError) {
			internal.Error(c, err)
			return
		} else if errno.Is(err, errno.LoginWrongInfoError) {
			internal.Error(c, err)
			return
		}
		
		internal.ServerError(c, errno.LoginServerError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			errors.WithStack(err), err.Error())
		return
	}
	
	str, err := pkg.GenerateToken(req.Account)
	if err != nil {
		internal.ServerError(c, errno.TokenGenerateError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			err)
		return
	}
	
	internal.Success(c, str)
}
