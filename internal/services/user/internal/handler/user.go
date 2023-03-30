package handler

import (
	"context"
	"gorm.io/gorm"
	"user/internal/repository"
	"user/internal/repository/model"
	CCNU "user/pkg"
	errno "user/pkg/errno"
	
	pb "user/proto"
)

type UserService struct {
	userDao repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		userDao: repository.NewUserRepository(db),
	}
}

func (u *UserService) UserRegister(ctx context.Context, in *pb.CCNUInfoRequest, resp *pb.Response) error {
	if err := CCNU.Login(in.Account, in.Password); err != nil {
		return err
	}
	
	if err := u.userDao.CreateUserIfNotExist(&model.User{
		Account:  in.Account,
		Password: in.Password,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Sex:      int32(in.Sex),
	}); err != nil {
		return errno.LoginServerError
	}
	
	return nil
}

func (u *UserService) UserLogin(ctx context.Context, in *pb.CCNUInfoRequest, resp *pb.CCNULoginResponse) error {
	return u.userDao.VerifyUser(in.Account, in.Password)
}

func (u *UserService) GetUserProfile(ctx context.Context, in *pb.Request, resp *pb.InnUserProfileResponse) error {
	//account, exist := ctx.Value("account").(string)
	//if !exist {
	//	return nil
	//}
	
	if one, err := u.userDao.GetUserProfile(in.Account); err != nil {
		return err
	} else {
		resp.Account = one.Account
		resp.Nickname = one.Nickname
		resp.Sex = pb.Sex(one.Sex)
		resp.Avatar = one.Avatar
	}
	return nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, in *pb.InnUserProfileRequest, resp *pb.Response) error {
	//account, exist := ctx.Value("account").(string)
	//if !exist {
	//	return nil
	//}
	var one = model.User{
		Account:  in.Account,
		Nickname: in.Nickname,
		Sex:      int32(in.Sex),
		Avatar:   in.Avatar,
	}
	if err := u.userDao.UpdateExistUserProfile(&one); err != nil {
		return err
	}
	return nil
}
