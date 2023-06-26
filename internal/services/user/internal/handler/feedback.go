package handler

import (
	"context"
	"user/internal/repository/model"
	pb "user/proto"
)

func (u *UserService) MakeFeedback(ctx context.Context, in *pb.MakeFeedbackRequest, resp *pb.Response) error {
	err := u.userDao.MakeFeedback(&model.Feedback{
		Account:      in.Account,
		PostId:       in.PostId,
		FeedbackType: in.FeedbackType,
		Content:      in.Content,
	})
	
	return err
}
