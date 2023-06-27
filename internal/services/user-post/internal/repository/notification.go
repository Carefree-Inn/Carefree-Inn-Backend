package repository

import (
	"github.com/pkg/errors"
	"user-post/internal/repository/model"
)

func (up *UserPost) GetNotificationHistory(account string, page, limit uint32) ([]*model.Notification, error) {
	data := make([]*model.Notification, 0, limit)
	
	if err := up.db.Table("notification").Where("to_user_account=?", account).
		Offset(int(page-1) * (int(limit))).Limit(int(limit)).Order("action_time").Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return data, nil
}
