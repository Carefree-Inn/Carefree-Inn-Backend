package model

import "time"

type PostLike struct {
	LikeId     uint32    `gorm:"column:like_id;primaryKey;autoIncrement"`
	Account    string    `gorm:"column:account;type:varchar(20);index"`
	PostId     uint32    `gorm:"column:post_id;index"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

type Comment struct {
	CommentId    uint32    `gorm:"column:comment_id;primaryKey;autoIncrement"`
	PostId       uint32    `gorm:"column:post_id;index"`
	FromUserId   string    `gorm:"column:from_id;type:varchar(20);index"`
	ToUserId     string    `gorm:"column:to_id;type:varchar(20);index"`
	Content      string    `gorm:"column:content;type:varchar(250)"`
	IsTop        bool      `gorm:"column:is_top;type:boolean"`
	TopCommentId uint32    `gorm:"column:top_comment_id"`
	CreateTime   time.Time `gorm:"column:create_time;autoCreateTime"`
}

type Notification struct {
	ActionType       string `gorm:"column:action_type;primaryKey"`
	FromUserAccount  string `gorm:"column:from_user_account"`
	FromUserNickname string `gorm:"column:from_user_nickname"`
	FromUserAvatar   string `gorm:"column:from_user_avatar"`
	ToUserAccount    string `gorm:"column:to_user_account"`
	PostId           uint32 `gorm:"post_id"`
	
	ActionId       uint32    `gorm:"column:action_id;primaryKey"`
	ActionTime     time.Time `gorm:"column:action_time"`
	CommentContent string    `gorm:"column:comment_content;type:varchar(250);default:null"`
}

func (PostLike) Table() string {
	return "post_like"
}

func (Comment) Table() string {
	return "comment"
}
