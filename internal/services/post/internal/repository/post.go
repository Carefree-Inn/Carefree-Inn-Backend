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

func (p *Post) GetLiked(account string, posts ...*model.Post) ([]*model.Post, error) {
	if account == "" {
		return posts, nil
	}
	
	ids := make([]uint32, 0, len(posts))
	for _, v := range posts {
		ids = append(ids, v.PostId)
	}
	
	var data = make([]uint32, 0, len(ids))
	if err := p.db.Table("post_like").Where("account = ? AND post_id IN (?)",
		account, ids).Pluck("post_id", &data).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	var hash = make(map[uint32]struct{})
	for _, val := range data {
		hash[val] = struct{}{}
	}
	
	for key, val := range posts {
		if _, ok := hash[val.PostId]; ok {
			posts[key].Liked = true
		}
	}
	
	return posts, nil
}

func (p *Post) GetPostOfUser(account string, page, limit int32) ([]*model.Post, error) {
	var posts = make([]*model.Post, 0, limit)
	if err := p.db.Table(model.Post{}.Table()).Where("account=?", account).
		Offset(int(page-1) * int(limit)).Limit(int(limit)).Preload("Tags").Order("create_time desc").
		Find(&posts).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return p.GetLiked(account, posts...)
}

func (p *Post) GetPostOfUserLiked(account string, page, limit int32) ([]*model.Post, error) {
	var posts = make([]*model.Post, 0, limit)
	if err := p.db.Table("post").
		Joins("JOIN post_like ON post.post_id = post_like.post_id AND post_like.account = ?", account).
		Offset(int(page-1) * int(limit)).Limit(int(limit)).Preload("Tags").
		Find(&posts).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	for key := range posts {
		posts[key].Liked = true
	}
	return posts, nil
}

func (p *Post) PostSquare() ([]*model.Tag, error) {
	tags := make([]*model.Tag, 0, 10)
	if err := p.db.Table("tag").Select([]string{"title", "count(*) as total"}).Group("title").
		Order("total DESC").Limit(10).Find(&tags).
		Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	return tags, nil
}

func (p *Post) GetPost(postId uint32, account string) ([]*model.Post, error) {
	post := make([]*model.Post, 0, 1)
	if err := p.db.Table(model.Post{}.Table()).Where("post_id=?", postId).Preload("Tags").
		Find(&post).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	
	data, err := p.GetLiked(account, post...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	
	return data, nil
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
