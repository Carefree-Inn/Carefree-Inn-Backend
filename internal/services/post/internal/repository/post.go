package repository

import (
	"github.com/pkg/errors"
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

func (p *Post) DeletePost(post *model.Post) error {
	session := p.db.Begin()
	
	one := model.Post{}
	if err := session.Table(post.Table()).Where("post_id=?", post.PostId).
		First(&one).Error; err != nil {
		return errors.WithStack(err)
	}
	
	if one.Account != post.Account {
		return errno.NoPowerDeleteError
	}
	
	if err := session.Table(post.Table()).
		Select("Tags").
		Delete(post).Error; err != nil {
		session.Rollback()
		return errors.WithStack(err)
	}
	
	session.Commit()
	return nil
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
