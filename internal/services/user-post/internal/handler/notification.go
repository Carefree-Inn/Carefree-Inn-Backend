package handler

import (
	"context"
	pb "user-post/proto"
)

func (up *UserPostService) GetNotificationHistory(ctx context.Context, in *pb.GetNotificationRequest, resp *pb.GetNotificationResponse) error {
	data, err := up.userPostDao.GetNotificationHistory(in.Account, in.Page, in.Limit)
	if err != nil {
		return err
	}
	
	notifications := make([]*pb.Notification, 0, len(data))
	for _, val := range data {
		notifications = append(notifications, &pb.Notification{
			ActionType:       val.ActionType,
			FromUserAccount:  val.FromUserAccount,
			FromUserNickname: val.FromUserNickname,
			FromUserAvatar:   val.FromUserAvatar,
			ToUserAccount:    val.ToUserAccount,
			PostId:           val.PostId,
			ActionId:         val.ActionId,
			ActionTime:       val.ActionTime.Format("2006-01-02 15:04:02"),
			CommentContent:   val.CommentContent,
		})
	}
	
	resp.Notifications = notifications
	resp.Account = in.Account
	
	return nil
}
