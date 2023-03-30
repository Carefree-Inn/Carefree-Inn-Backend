package model

import "time"

type Like struct {
	LikeId     uint32    `gorm:"column:like_id;primaryKey"`
	Uuid       string    `gorm:"column:uuid"`
	PostId     uint32    `gorm:"column:post_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type Comment struct {
	CommentId    uint32    `gorm:"column:comment_id"`
	Uuid         [16]byte  `gorm:"column:uuid"`
	PostId       uint32    `gorm:"column:post_id"`
	ReplyComment uint32    `gorm:"column:reply_comment"`
	CreateTime   time.Time `gorm:"column:create_time"`
	DeletedTime  time.Time `gorm:"column:deleted_time"`
}

func (Like) Table() string {
	return "like"
}

func (Comment) Table() string {
	return "comment"
}
