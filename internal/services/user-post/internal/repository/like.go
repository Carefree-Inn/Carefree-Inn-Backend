package repository

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"user-post/internal/repository/model"
)

type LikeInfo struct {
	PostId     uint32 `json:"post_id"`
	Account    string `json:"account"`
	CreateTime string `json:"create_time"`
	LikeType   string `json:"like_type"`
}

func (l *LikeInfo) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *LikeInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}

func (up *UserPost) MakeLike(postId uint32, account string) error {
	info, err := (&LikeInfo{
		PostId:     postId,
		Account:    account,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		LikeType:   "make",
	}).Marshal()
	
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(
		up.rdb.Publish(context.TODO(), "like_before", info).Err(),
	)
}

func (up *UserPost) DeleteLike(postId uint32, account string) error {
	info, err := json.Marshal(&LikeInfo{
		PostId:     postId,
		Account:    account,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		LikeType:   "delete",
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(
		up.rdb.Publish(context.TODO(), "like_before ", info).Err(),
	)
}

func (up *UserPost) GetLikes(account string, page int32, limit int32) ([]*model.Like, error) {
	likes := make([]*model.Like, 20)
	if err := up.db.Table(model.Like{}.Table()).Where("account=?",
		account).Offset(int((page - 1) * limit)).Limit(int(limit)).Find(&likes).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return likes, nil
}

func (up *UserPost) IsBatchLiked(account string, postIds []uint32) ([]*model.Like, error) {
	likes := make([]*model.Like, 20)
	if err := up.db.Table(model.Like{}.Table()).Where("account=? AND post_id in ?",
		account, postIds).Find(&likes).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return likes, nil
}
