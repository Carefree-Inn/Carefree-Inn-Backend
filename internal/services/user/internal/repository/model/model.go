package model

import "time"

type User struct {
	Account    string    `gorm:"column:account;type:VARCHAR(20);primaryKey"`
	Password   string    `gorm:"column:password;type:VARCHAR(200)"`
	Nickname   string    `gorm:"column:nickname;type:VARCHAR(30)"`
	Sex        int32     `gorm:"column:sex"`
	Avatar     string    `gorm:"column:avatar;type:VARCHAR(200)"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

// 功能反馈
type Feedback struct {
	Account      string `gorm:"column:account;primaryKey""`
	FeedbackType string `gorm:"column:feedback_type;type:VARCHAR(100)"`
	Content      string `gorm:"content;type:TEXT"`
}

// 收藏
type Collection struct {
	Account string `gorm:"column:account;"`
}

func (Collection) Table() string {
	return "collection"
}

func (User) Table() string {
	return "user"
}

func (Feedback) Table() string {
	return "feedback"
}
