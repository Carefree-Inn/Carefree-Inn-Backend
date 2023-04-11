package handler

import (
	"context"
	"gorm.io/gorm"
	"post/internal/repository"
	"post/internal/repository/model"
	pb "post/proto"
)

type PostService struct {
	postDao repository.PostRepository
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		postDao: repository.NewPostRepository(db),
	}
}

func (p *PostService) CreatePost(ctx context.Context, in *pb.CreatePostRequest, resp *pb.Response) error {
	var tags = make([]*model.Tag, 0, 16)
	
	for _, v := range in.Tag {
		tags = append(tags, &model.Tag{
			Title: v.Title,
			TagId: v.TagId,
		})
	}
	
	var post = &model.Post{
		Title:      in.Title,
		CategoryId: in.Category.CategoryId,
		Content:    in.Content,
		Account:    in.Account,
		Star:       0,
		Comments:   0,
		Tags:       tags,
	}
	
	if err := p.postDao.CreatePost(post); err != nil {
		return err
	}
	
	return nil
}
func (p *PostService) DeletePost(ctx context.Context, in *pb.DeletePostRequest, resp *pb.Response) error {
	return p.postDao.DeletePost(&model.Post{
		PostId:  in.PostId,
		Account: in.Account,
	})
}
