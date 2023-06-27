package repository

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"user-post/internal/repository/model"
	"user-post/pkg/errno"
)

type LikeInfo struct {
	PostId uint32 `json:"post_id"`
	
	ToUserAccount string `json:"to_user_account"`
	CreateTime    string `json:"create_time"`
	LikeType      string `json:"like_type"`
	
	FromUserAccount  string `json:"from_user_account"`
	FromUserAvatar   string `json:"from_user_avatar"`
	FromUserNickname string `json:"from_user_nickname"`
}

func (l *LikeInfo) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *LikeInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}

func (up *UserPost) MakeLike(likeInfo *LikeInfo) error {
	info, err := likeInfo.Marshal()
	
	if err != nil {
		return errors.WithStack(err)
	}
	
	like := model.PostLike{}
	if err := up.db.Table(like.Table()).Where("post_id = ? AND account = ?", likeInfo.PostId, likeInfo.FromUserAccount).
		Find(&like).Error; err != nil {
		return errors.WithStack(err)
	}
	
	if like.Account == likeInfo.ToUserAccount {
		return errno.DuplicateLike
	}
	
	return errors.WithStack(
		up.rdb.Publish(context.TODO(), "like", info).Err(),
	)
}

func (up *UserPost) DeleteLike(postId uint32, account string) error {
	info, err := (&LikeInfo{
		PostId:          postId,
		FromUserAccount: account,
		CreateTime:      time.Now().Format("2006-01-02 15:04:05"),
		LikeType:        "delete",
	}).Marshal()
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(
		up.rdb.Publish(context.TODO(), "like", info).Err(),
	)
}

func (up *UserPost) GetLikes(account string, page int32, limit int32) ([]*model.PostLike, error) {
	likes := make([]*model.PostLike, 20)
	if err := up.db.Table(model.PostLike{}.Table()).Where("account=?",
		account).Offset(int((page - 1) * limit)).Limit(int(limit)).Order("create_time desc").Find(&likes).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return likes, nil
}

func (up *UserPost) IsBatchLiked(account string, postIds []uint32) ([]*model.PostLike, error) {
	likes := make([]*model.PostLike, 20)
	if err := up.db.Table(model.PostLike{}.Table()).Where("account=? AND post_id in ?",
		account, postIds).Find(&likes).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return likes, nil
}
