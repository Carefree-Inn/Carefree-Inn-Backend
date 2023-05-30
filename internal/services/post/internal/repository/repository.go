package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"post/config"
	"post/internal/repository/model"
	"post/pkg/log"
	"strings"
	"sync"
)

type PostRepository interface {
	CreatePost(post *model.Post) error
	GetCategory(id []uint32) ([]*model.Category, error)
	DeletePost(post *model.Post) error
	GetAllCategory() ([]*model.Category, error)
	GetPostOfTag(title string, account string) ([]*model.Post, error)
	GetPostOfCategory(category *model.Category, account string, page uint32, limit uint32) ([]*model.Post, error)
	SearchPost(content string, searchType string, account string) ([]*model.Post, error)
	GetPostOfUser(account string, page int32, limit int32) ([]*model.Post, error)
	GetPostOfUserLiked(account string, page, limit int32) ([]*model.Post, error)
}

func NewPostRepository(dbUp *gorm.DB) PostRepository {
	var one = new(Post)
	one.init(dbUp)
	return one
}

type Post struct {
	db   *gorm.DB
	once sync.Once
}

func (d *Post) init(dbUp *gorm.DB) {
	if d.db == nil {
		d.once.Do(func() {
			d.db = dbUp
		})
	}
}

func Init(cfg config.Config) *gorm.DB {
	Db, err := gorm.Open(mysql.Open(cfg.Database.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据库初始化失败")
	}
	
	if err := Db.AutoMigrate(&model.Post{}, &model.Category{},
		&model.Tag{}); err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据表初始化失败")
	}
	
	if cfg.Category.Start {
		if err := Db.Table(model.Category{}.Table()).
			Exec(insertBatchCategory(cfg.Category.Region)).Error;
			err != nil {
			log.Fatal(nil, errors.WithStack(err), "分区初始化失败")
		}
	}
	return Db
}

func insertBatchCategory(data []string) string {
	var sb strings.Builder
	sb.WriteString("INSERT INTO `category`(`title`) VALUES")
	for k, v := range data {
		if k != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(fmt.Sprintf(`("%s")`, v))
	}
	sb.WriteByte(';')
	return sb.String()
}
