package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"post/internal/repository/model"
	"strings"
)

func (p *Post) SearchPost(content string, searchType string, account string) ([]*model.Post, error) {
	var str strings.Builder
	str.WriteByte('%')
	for _, v := range content {
		str.WriteRune(v)
		str.WriteByte('%')
	}
	var posts = make([]*model.Post, 0, 100)
	if searchType == "title" {
		if err := p.db.Table(model.Post{}.Table()).Preload("Tags").Where(
			fmt.Sprintf("title LIKE '%s'", str.String())).
			Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if searchType == "content" {
		if err := p.db.Table(model.Post{}.Table()).Where(
			fmt.Sprintf("content LIKE %s", str.String())).
			Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if searchType == "all" {
		if err := p.db.Table(model.Post{}.Table()).Where(
			fmt.Sprintf("content LIKE %s OR title LIKE %s", str.String(), str.String())).
			Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return p.GetLiked(posts, account)
}
