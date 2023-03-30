package model

import "time"

type Category struct {
	CategoryId uint32 `gorm:"column:category_id"`
	Title      string `gorm:"column:title;type:VARCHAR(200)"`
}

type Post struct {
	PostId      uint32    `gorm:"column:post_id;primaryKey"`
	Uuid        string    `gorm:"column:uuid;type:VARCHAR(20)"`
	CategoryId  uint32    `gorm:"column:category_id"`
	Title       string    `gorm:"column:title;type:VARCHAR(200)"`
	Content     string    `gorm:"column:content;type:TEXT"`
	CreateTime  time.Time `gorm:"column:create_time"`
	DeletedTime time.Time `gorm:"column:deleted_time"`
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

type Tag struct {
	TagId uint32 `gorm:"column:tag_id"`
	Title string `gorm:"column:title;type:VARCHAR(200)"`
}

type PostTag struct {
	PostId uint32 `gorm:"column:post_id"`
	TagId  uint32 `gorm:"column:tag_id"`
}

func (Post) Table() string {
	return "`post`"
}

func (PostTag) Table() string {
	return "`post_tag`"
}

func (Tag) Table() string {
	return "`tag`"
}
