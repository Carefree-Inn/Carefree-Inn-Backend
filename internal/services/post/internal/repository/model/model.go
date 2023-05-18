package model

import (
	"time"
)

type Category struct {
	CategoryId uint32 `gorm:"column:category_id;primaryKey;autoIncrement"`
	Title      string `gorm:"column:title;type:VARCHAR(200);unique"`
}

type Post struct {
	PostId     uint32    `gorm:"column:post_id;primaryKey;autoIncrement"`
	Account    string    `gorm:"column:account;type:VARCHAR(20)"`
	CategoryId uint32    `gorm:"column:category_id"`
	Title      string    `gorm:"column:title;type:VARCHAR(200)"`
	Content    string    `gorm:"column:content;type:TEXT"`
	Likes      uint32    `gorm:"column:likes"`
	Comments   uint32    `gorm:"column:comments"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	Tags       []*Tag    `gorm:"many2many:post_tags;foreignKey:post_id"`
}

type Tag struct {
	TagId uint32  `gorm:"column:tag_id;primaryKey;autoIncrement"`
	Title string  `gorm:"column:title;type:VARCHAR(200)"`
	Posts []*Post `gorm:"many2many:post_tags;foreignKey:tag_id"`
}

func (Post) Table() string {
	return "`post`"
}

func (Tag) Table() string {
	return "`tag`"
}

func (Category) Table() string {
	return "`category`"
}

//type Feedback struct {
//	FeedbackId   uint32 `gorm:"column:feedback_id"`
//	PostId       uint32 `gorm:"column:post_id"`
//	FeedbackType string `gorm:"column:feedback_type;type:VARCHAR(200)"`
//	Content      string `gorm:"column:content;type:TEXT"`
//}
//
//type PostFeedback struct {
//
//}
