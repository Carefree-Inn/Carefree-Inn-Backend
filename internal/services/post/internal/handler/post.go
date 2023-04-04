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
		CategoryId: in.CategoryId,
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
func (p *PostService) GetPostOfCategory(ctx context.Context, in *pb.CategoryRequest, resp *pb.PostResponse) error {
	posts, err := p.postDao.GetPostOfCategory(&model.Category{
		CategoryId: in.Category.CategoryId,
		Title:      in.Category.Title,
	}, in.Page, in.Limit)
	if err != nil {
		return err
	}
	
	var data = make([]*pb.PostInfo, 0, len(posts))
	for _, v := range posts {
		var tags = make([]*pb.TagInfo, 0, len(v.Tags))
		for _, val := range v.Tags {
			tags = append(tags, &pb.TagInfo{
				Title: val.Title,
				TagId: val.TagId,
			})
		}
		data = append(data, &pb.PostInfo{
			Account:    v.Account,
			PostId:     v.PostId,
			Title:      v.Title,
			Content:    v.Content,
			CategoryId: v.CategoryId,
			Star:       v.Star,
			Comments:   v.Comments,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
			Tag:        tags,
		})
	}
	resp.Posts = data
	return nil
}
