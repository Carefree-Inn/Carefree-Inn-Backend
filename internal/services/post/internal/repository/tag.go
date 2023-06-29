package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"post/internal/repository/model"
	"time"
)

func CheckTag(db *gorm.DB, titles []string) ([]*model.Tag, []*model.Tag, error) {
	var exist = make([]*model.Tag, 0, len(titles))
	// 存在的tag
	if err := db.Table(model.Tag{}.Table()).Where("title IN ?", titles).
		Find(&exist).Error; err != nil {
		return nil, nil, errors.WithStack(err)
	}
	// 不存在的tag
	var notExist = make([]*model.Tag, 0, len(titles))
	for _, v := range titles {
		in := false
		for _, val := range exist {
			if v == val.Title {
				in = true
				break
			}
		}
		if !in {
			notExist = append(notExist, &model.Tag{
				Title: v,
			})
		}
	}
	return notExist, exist, nil
}

func (p *Post) GetPostOfTag(title string, account string) ([]*model.Post, error) {
	var data = make([]*model.Post, 0, 16)
	minute := time.Now().Minute() / 5
	if err := p.db.Table(model.Post{}.Table()).
		Joins("JOIN post_tags pt ON post.post_id = pt.post_post_id").
		Joins("JOIN tag t ON t.tag_id = pt.tag_tag_id").
		Where("t.title = ?", title).Order(fmt.Sprintf("RAND(%d)", minute)).
		Preload("Tags").
		Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return p.GetLiked(account, data...)
}
