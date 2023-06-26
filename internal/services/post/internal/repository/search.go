package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"post/internal/repository/model"
	"strings"
	"time"
)

func (p *Post) SearchPost(content string, searchType string, account string) ([]*model.Post, error) {
	var str strings.Builder
	str.WriteByte('%')
	for _, v := range content {
		str.WriteRune(v)
		str.WriteByte('%')
	}
	
	data := str.String()
	minute := time.Now().Minute() / 5
	var posts = make([]*model.Post, 0, 100)
	if searchType == "title" {
		if err := p.db.Table(model.Post{}.Table()).Preload("Tags").Where(
			fmt.Sprintf("title LIKE '%s'", data)).Order(fmt.Sprintf("RAND(%d)", minute)).
			Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if searchType == "content" {
		if err := p.db.Table(model.Post{}.Table()).Where(
			fmt.Sprintf("content REGEXP '%s'", fmt.Sprintf(".*>[^<]%s[^<]<\\/Text>", strings.Replace(data, "%", "*", -1)))).
			Order(fmt.Sprintf("RAND(%d)", minute)).Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if searchType == "all" {
		if err := p.db.Table(model.Post{}.Table()).Where(
			fmt.Sprintf("content REGEXP '%s' OR title LIKE '%s'", fmt.Sprintf(".*>[^<]%s[^<]*<\\/Text>", strings.Replace(data, "%", "*", -1)), data)).
			Order(fmt.Sprintf("RAND(%d)", minute)).Find(&posts).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return p.GetLiked(posts, account)
}
