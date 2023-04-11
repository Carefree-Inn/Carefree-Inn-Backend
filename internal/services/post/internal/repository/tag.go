package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"post/internal/repository/model"
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

func (p *Post) GetPostOfTag(title string) ([]*model.Post, error) {
	var data = make([]*model.Post, 0, 16)
	
	if err := p.db.Table(model.Post{}.Table()).Preload("Tags").Exec("SELECT * FROM `post` WHERE `post_id` IN" +
		"(SELECT `post_post_id` FROM `post_tags` JOIN `tag` ON `tag`.`tag_id`=" +
		"`post_tags`.`tag_tag_id`)").Find(&data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return data, nil
}
