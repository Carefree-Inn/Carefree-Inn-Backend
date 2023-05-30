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
	
	_, err := u.UserLogin(ctx, &pb.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		if errno.Is(err, errno.LoginWrongInfoError) {
			internal.Error(c, err)
			return
		}
		
		internal.ServerError(c, errno.LoginServerError.Error())
		return
	}
	
	str, err := pkg.GenerateToken(req.Account)
	if err != nil {
		internal.ServerError(c, errno.TokenGenerateError.Error())
		return
	}
	
	internal.Success(c, str)
}
