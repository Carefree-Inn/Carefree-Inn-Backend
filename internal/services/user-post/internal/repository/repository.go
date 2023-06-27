package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
	"user-post/internal/repository/model"
	"user-post/pkg/log"
)

func Init(dsn string) *Database {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据库初始化失败")
	}
	
	if err := db.AutoMigrate(&model.PostLike{}, &model.Comment{}, &model.Notification{}); err != nil {
		log.Fatal(nil, errors.WithStack(err), "数据表初始化失败")
	}
	
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	for err := rdb.Ping(context.TODO()).Err(); err != nil; {
		err = rdb.Ping(context.TODO()).Err()
		time.Sleep(time.Millisecond * 500)
	}
	
	return &Database{
		Db:  db,
		Rdb: rdb,
	}
}

type Database struct {
	Rdb *redis.Client
	Db  *gorm.DB
}

type UserPost struct {
	db   *gorm.DB
	rdb  *redis.Client
	once sync.Once
}

func (d *UserPost) init(database *Database) {
	if d.db == nil {
		d.once.Do(func() {
			d.db = database.Db
			d.rdb = database.Rdb
			go d.processLike()
		})
	}
}

func (d *UserPost) processLike() {
	like := d.rdb.Subscribe(context.Background(), "like")
	likeCh := like.Channel()
	
	for msg := range likeCh {
		print("in")
		likeInfo := LikeInfo{}
		err := json.Unmarshal([]byte(msg.Payload), &likeInfo)
		if err != nil {
			log.Warn(nil, err, "订阅者消息序列化失败")
			continue
		}
		session := d.db.Begin()
		fmt.Println(likeInfo)
		switch likeInfo.LikeType {
		case "make":
			createTime, err := time.Parse("2006-01-02 15:04:05", likeInfo.CreateTime)
			if err != nil {
				createTime = time.Now()
			}
			postLike := model.PostLike{
				PostId:     likeInfo.PostId,
				Account:    likeInfo.FromUserAccount,
				CreateTime: createTime,
			}
			if err := session.Table(model.PostLike{}.Table()).Create(
				&postLike).Error; err != nil {
				log.Warn(log.WithFields(logrus.Fields{
					"post_id": likeInfo.PostId,
					"account": likeInfo.FromUserAccount,
				}), errors.WithStack(err), "点赞失败")
				session.Rollback()
				continue
			} else {
				if errAdd := session.Exec("UPDATE post SET likes = likes + 1 WHERE post_id = ?",
					likeInfo.PostId).Error; errAdd != nil {
					log.Warn(log.WithFields(logrus.Fields{
						"post_id": likeInfo.PostId,
						"account": likeInfo.FromUserAccount,
					}), errors.WithStack(err), "点赞失败")
					session.Rollback()
					continue
				} else {
					if errNotification := d.SetNotification(&likeInfo, postLike.LikeId, postLike.CreateTime); errNotification != nil {
						log.Warn(log.WithFields(logrus.Fields{
							"post_id": likeInfo.PostId,
							"account": likeInfo.FromUserAccount,
						}), errors.WithStack(err), "点赞失败")
						session.Rollback()
						continue
					}
				}
				
				log.Info(log.WithFields(logrus.Fields{
					"post_id": likeInfo.PostId,
					"account": likeInfo.FromUserAccount,
				}), "点赞成功")
				session.Commit()
				
				d.rdb.Publish(context.TODO(), "like_after", msg.Payload)
			}
		
		case "delete":
			if err := session.Where(
				"post_id=? AND account=?", likeInfo.PostId,
				likeInfo.FromUserAccount).Delete(&model.PostLike{}).Error; err != nil {
				log.Warn(log.WithFields(logrus.Fields{
					"post_id": likeInfo.PostId,
					"account": likeInfo.FromUserAccount,
				}), errors.WithStack(err), "取消点赞失败")
				session.Rollback()
				continue
			} else {
				if errAdd := session.Exec("UPDATE post SET likes = likes - 1 WHERE post_id = ?",
					likeInfo.PostId).Error; errAdd != nil {
					log.Warn(log.WithFields(logrus.Fields{
						"post_id": likeInfo.PostId,
						"account": likeInfo.FromUserAccount,
					}), errors.WithStack(err), "点赞失败")
					session.Rollback()
					continue
				}
				
				log.Info(log.WithFields(logrus.Fields{
					"post_id": likeInfo.PostId,
					"account": likeInfo.FromUserAccount,
				}), "取消点赞成功")
				session.Commit()
				
			}
		}
	}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type UserPostRepository interface {
	DeleteComment(comment *model.Comment) error
	MakeComment(comment *model.Comment, commentInfo *Comment) error
	GetCommentOfPost(postId uint32, page, limit uint32) ([]*model.Comment, error)
	GetCommentOfUser(account string, page, limit uint32) ([]*model.Comment, error)
	
	MakeLike(likeInfo *LikeInfo) error
	DeleteLike(postId uint32, account string) error
	GetLikes(account string, page int32, limit int32) ([]*model.PostLike, error)
	IsBatchLiked(account string, postIds []uint32) ([]*model.PostLike, error)
	
	GetNotificationHistory(account string, page, limit uint32) ([]*model.Notification, error)
}

func NewUserPostRepository(database *Database) UserPostRepository {
	var one = new(UserPost)
	one.init(database)
	return one
}
