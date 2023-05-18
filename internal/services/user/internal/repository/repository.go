package repository

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"user/internal/repository/model"
	"user/pkg/log"
)

func Init(dsn string) *gorm.DB {
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据库初始化失败")
	}
	
	if err := Db.AutoMigrate(&model.User{}, &model.Feedback{}); err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据表初始化失败")
	}
	return Db
}

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

type UserRepository interface {
	UpdateExistUserProfile(user *model.User) error
	CreateUser(user *model.User) error
	GetUserProfile(account string) (*model.User, error)
	VerifyUser(account, password string) error
	GetBatchUserProfile(accounts []string) ([]*model.User, error)
}

func NewUserRepository(dbUp *gorm.DB) UserRepository {
	var one = new(User)
	one.init(dbUp)
	return one
}
