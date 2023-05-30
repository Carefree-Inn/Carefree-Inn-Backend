package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"post/internal/repository/model"
	"post/pkg/errno"
)

func (p *Post) GetCategory(id []uint32) ([]*model.Category, error) {
	var category = make([]*model.Category, 0)
	if err := p.db.Table(model.Category{}.Table()).Where("category_id IN ?", id).
		Find(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.CategoryNotExistError
		}
		return nil, errors.WithStack(err)
	}
	return category, nil
}

func (p *Post) GetAllCategory() ([]*model.Category, error) {
	var categories = make([]*model.Category, 0, 16)
	if err := p.db.Table(model.Category{}.Table()).
		Find(&categories).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return categories, nil
}

func (p *Post) GetPostOfCategory(category *model.Category, account string, page, limit uint32) ([]*model.Post, error) {
	var posts = make([]*model.Post, 0, 16)
	if err := p.db.Table(model.Post{}.Table()).
		Where("category_id=?", category.CategoryId).Preload("Tags").
		Offset(int((page - 1) * limit)).Limit(int(limit)).
		Find(&posts).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return p.GetLiked(posts, account)
}
