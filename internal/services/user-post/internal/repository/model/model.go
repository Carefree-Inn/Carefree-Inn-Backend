package model

import "time"

type Like struct {
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

func (Like) Table() string {
	return "like"
}

func (Comment) Table() string {
	return "comment"
}
