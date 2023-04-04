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
	session := d.db.Begin()
	if errors.Is(d.ExistUser(user.Account), gorm.ErrRecordNotFound) {
		if err := session.Table(user.Table()).Create(&user).Error; err != nil {
			session.Rollback()
			return errors.WithStack(err)
		}
	}
	return errors.WithStack(session.Commit().Error)
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
	session := d.db.Begin()
	if err := session.Table(user.Table()).Updates(&user).Error;
		err != nil {
		session.Rollback()
		return errors.WithStack(err)
	}
	return errors.WithStack(session.Commit().Error)
}

func (d *User) GetBatchUserProfile(accounts []string) ([]*model.User, error) {
	var users = make([]*model.User, 0, len(accounts))
	if err := d.db.Table(model.User{}.Table()).Where("account IN ?", accounts).
		Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}
