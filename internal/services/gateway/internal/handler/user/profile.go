package user

import (
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	pb "github.com/jackj-ohn1/package/proto/user"
	"golang.org/x/net/context"
)

type UserInfo struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Sex      int8   `json:"sex"`
	Avatar   string `json:"avatar"`
}

//  GetProfile getProfile
//	@Summary		获取用户信息 api
//	@Tags			user
//	@Description	获取用户信息
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//	@Success		200				{object}	internal.Response
//	@Router			/user/profile [get]
func (u *userHandler) GetProfile(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	profile, err := u.GetUserProfile(ctx, &pb.Request{
		Account: account,
	})
	if err != nil {
		internal.ServerError(c, errno.DatabaseError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			err)
		return
	}
	
	internal.Success(c, UserInfo{
		Account:  profile.Account,
		Nickname: profile.Nickname,
		Sex:      int8(profile.Sex),
		Avatar:   profile.Avatar,
	})
}

//  UpdateProfile updateProfile
//	@Summary		修改用户信息 api
//	@Tags			user
//	@Description	修改用户信息
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string		true	"用户token"
//	@Param			object			body		UserInfo	true	"需要修改的信息"
//	@Success		200				{object}	internal.Response
//	@Router			/user/profile [put]
func (u *userHandler) UpdateProfile(c *gin.Context) {
	var req UserInfo
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		internal.Error(c, errno.JsonDataError)
		log.Info(
			log.WithField("X-Request-Id", c.MustGet("uuid")),
			errno.JsonDataError.Error(),
		)
		return
	}
	
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	account := c.MustGet("account").(string)
	_, err := u.UpdateUserProfile(ctx, &pb.InnUserProfileRequest{
		Account:  account,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Sex:      pb.Sex(req.Sex),
	})
	if err != nil {
		internal.ServerError(c, errno.DatabaseError.Error())
		log.Panic(log.WithField("X-Request-Id", c.MustGet("uuid")),
			err)
		return
	}
	
	internal.Success(c, nil)
}
