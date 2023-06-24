package handler

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"user/internal/repository"
	"user/internal/repository/model"
	CCNU "user/pkg"
	"user/pkg/errno"
	
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

func (u *UserService) UserLogin(ctx context.Context, in *pb.LoginRequest, resp *pb.LoginResponse) error {
	verifyErr := u.userDao.VerifyUser(in.Account, in.Password)
	if verifyErr != nil {
		if errors.Is(verifyErr, errno.UserNotExistError) {
			if err := CCNU.Login(in.Account, in.Password); err != nil {
				return errno.LoginWrongInfoError
			}
			if err := u.userDao.CreateUser(&model.User{
				Account:  in.Account,
				Password: in.Password,
				Nickname: in.Account,
				Avatar:   CCNU.DefaultAvatat,
			}); err != nil {
				return err
			}
		}
		return verifyErr
	}
	return nil
}

func (u *UserService) GetUserProfile(ctx context.Context, in *pb.Request, resp *pb.UserProfileResponse) error {
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

func (u *UserService) UpdateUserProfile(ctx context.Context, in *pb.UserProfileRequest, resp *pb.Response) error {
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

func (u *UserService) GetBatchUserProfile(ctx context.Context, in *pb.BatchUserProfileRequest, resp *pb.BatchUserProfileResponse) error {
	data, err := u.userDao.GetBatchUserProfile(in.Accounts)
	if err != nil {
		return err
	}
	
	resp.Data = make(map[string]*pb.UserProfileResponse)
	for _, v := range data {
		resp.Data[v.Account] = &pb.UserProfileResponse{
			Account:  v.Account,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Sex:      pb.Sex(v.Sex),
		}
	}
	
	return nil
}
