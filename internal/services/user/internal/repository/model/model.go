package model

import "time"

type User struct {
	Account    string    `gorm:"column:account;type:VARCHAR(20);primaryKey"`
	Password   string    `gorm:"column:password;type:VARCHAR(20)"`
	Nickname   string    `gorm:"column:nickname;type:VARCHAR(30)"`
	Sex        int32     `gorm:"column:sex"`
	Avatar     string    `gorm:"column:avatar;type:VARCHAR(200)"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

// 功能反馈
type Feedback struct {
	Account string `gorm:"column:account;primaryKey""`
	Content string `gorm:"content"`
}

// 点赞
type Star struct {
}

// 收藏
type Collection struct {
	Account string `gorm:"column:account;"`
}

func (Star) Table() string {
	return "star"
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
