package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"post/internal/repository/model"
	"post/pkg/errno"
)

func (p *Post) CreatePost(post *model.Post) error {
	session := p.db.Begin()
	if err := session.Table(post.Table()).Create(post).
		Error; err != nil {
		session.Rollback()
		return errors.WithStack(err)
	}
	
	return errors.WithStack(session.Commit().Error)
}

//func (p *Post) CreatePost(post *model.Post) error {
//	session := p.db.Begin()
//
//	// 创建tag
//	if err := session.Table(post.Table()).Create(post).Error; err != nil {
//		session.Rollback()
//		return errors.WithStack(err)
//	}
//
//	var tags = make([]string, 0, len(post.Tags))
//	for _, v := range post.Tags {
//		tags = append(tags, v.Title)
//	}
//
//	// 创建tag
//	notExist, exist, err := CheckTag(session, tags)
//	if err != nil {
//		session.Rollback()
//		return err
//	} else if len(notExist) > 0 {
//		if err := session.Table(model.PostTag{}.Table()).Create(notExist).
//			Error; err != nil {
//			session.Rollback()
//			return errors.WithStack(err)
//		}
//	}
//
//	// 创建post_tag
//	var postTags = make([]*model.PostTag, 0, len(notExist)+len(exist))
//	for _, v := range notExist {
//		postTags = append(postTags, &model.PostTag{
//			PostId: post.PostId,
//			TagId:  v.TagId,
//		})
//	}
//	for _, v := range exist {
//		postTags = append(postTags, &model.PostTag{
//			PostId: post.PostId,
//			TagId:  v.TagId,
//		})
//	}
//	if err := session.Table(model.PostTag{}.Table()).Create(postTags).
//		Error; err != nil {
//		session.Rollback()
//		return errors.WithStack(err)
//	}
//	return errors.WithStack(session.Commit().Error)
//}

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

func (p *Post) GetCategory(title string) (*model.Category, error) {
	var category model.Category
	if err := p.db.Table(category.Table()).Where("title=?", title).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.CategoryNotExistError
		}
		return nil, errors.WithStack(err)
	}
	return &category, nil
}

func (p *Post) GetAllCategory() ([]*model.Category, error) {
	var categories = make([]*model.Category, 0, 16)
	if err := p.db.Table(model.Category{}.Table()).
		Find(&categories).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return categories, nil
}

func (p *Post) GetPostOfCategory(category *model.Category, page, limit uint32) ([]*model.Post, error) {
	var posts = make([]*model.Post, 0, 16)
	if err := p.db.Table(model.Post{}.Table()).
		Where("category_id=?", category.CategoryId).Preload("Tags").
		Offset(int((page - 1) * limit)).Limit(int(limit)).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *Post) DeletePost(post *model.Post) error {
	session := p.db.Begin()
	
	if err := session.Select("tags").
		Delete(post).Error; err != nil {
		session.Rollback()
		return errors.WithStack(err)
	}
	
	session.Commit()
	return nil
}
