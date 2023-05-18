package handler

import (
	"context"
	"post/internal/repository/model"
	pb "post/proto"
)

func (p *PostService) GetCategory(ctx context.Context, in *pb.Request, resp *pb.CategoryResponse) error {
	categories, err := p.postDao.GetAllCategory()
	if err != nil {
		return err
	}
	
	var data = make([]*pb.CategoryInfo, 0, len(categories))
	for _, v := range categories {
		data = append(data, &pb.CategoryInfo{
			Title:      v.Title,
			CategoryId: v.CategoryId,
		})
	}
	resp.Categories = data
	return nil
}

func (p *PostService) GetPostOfCategory(ctx context.Context, in *pb.PostOfCategoryRequest, resp *pb.PostResponse) error {
	posts, err := p.postDao.GetPostOfCategory(&model.Category{
		CategoryId: in.Category.CategoryId,
		Title:      in.Category.Title,
	}, in.Page, in.Limit)
	if err != nil {
		return err
	}
	
	data, err := p.convertPost(posts)
	if err != nil {
		return err
	}
	
	resp.Posts = data
	return nil
}
