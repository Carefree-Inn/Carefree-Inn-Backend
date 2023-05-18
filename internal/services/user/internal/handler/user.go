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

func (u *UserService) UserLogin(ctx context.Context, in *pb.CCNUInfoRequest, resp *pb.CCNULoginResponse) error {
	verifyErr := u.userDao.VerifyUser(in.Account, in.Password)
	if verifyErr != nil {
		if errors.Is(verifyErr, errno.UserNotExistError) {
			if err := u.userDao.CreateUser(&model.User{
				Account:  in.Account,
				Password: in.Password,
				Nickname: in.Account,
				Avatar:   CCNU.GetAvatar(),
			}); err != nil {
				return err
			}
		}
		return verifyErr
	}
	return nil
}

func (u *UserService) GetUserProfile(ctx context.Context, in *pb.Request, resp *pb.InnUserProfileResponse) error {
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

func (u *UserService) GetBatchUserProfile(ctx context.Context, in *pb.BatchUserProfileRequest, resp *pb.BatchUserProfileResponse) error {
	data, err := u.userDao.GetBatchUserProfile(in.Accounts)
	if err != nil {
		return err
	}
	
	var users = make([]*pb.InnUserProfileResponse, 0, len(data))
	for _, v := range data {
		users = append(users, &pb.InnUserProfileResponse{
			Account:  v.Account,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Sex:      pb.Sex(v.Sex),
		})
	}
	
	resp.Data = users
	return nil
}
