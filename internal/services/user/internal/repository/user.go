package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
	"user/internal/repository/model"
	"user/pkg/errno"
)

type User struct {
	db   *gorm.DB
	once sync.Once
}

func (d *User) init(dbUp *gorm.DB) {
	if d.db == nil {
		d.once.Do(func() {
			d.db = dbUp
		})
	}
}

func (d *User) ExistUser(account string) error {
	var user model.User
	return d.db.Table(user.Table()).Where("account=?", account).First(&user).Error
}

func (d *User) VerifyUser(account, password string) error {
	var user model.User
	if err := d.db.Table(user.Table()).Where("account=?", account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.UserNotExistError
		}
		return errors.WithStack(err)
	}
	if user.Password != password {
		return errno.LoginWrongInfoError
	}
	return nil
}

func (d *User) CreateUserIfNotExist(user *model.User) error {
	if errors.Is(d.ExistUser(user.Account), gorm.ErrRecordNotFound) {
		return errors.WithStack(d.db.Table(user.Table()).Create(&user).Error)
	}
	return nil
}

func (d *User) GetUserProfile(account string) (*model.User, error) {
	var one model.User
	if err := d.db.Table(one.Table()).Omit("password").
		Where("account=?", account).Find(&one).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &one, nil
}

func (d *User) UpdateExistUserProfile(user *model.User) error {
	return errors.WithStack(d.db.Table(user.Table()).Updates(&user).Error)
}
