package handler

import (
	"context"
	"github.com/pkg/errors"
	"post/pkg/errno"
	pb "post/proto"
)

func (p *PostService) SearchPost(ctx context.Context, req *pb.SearchRequest, resp *pb.PostResponse) error {
	if req.Content == "" || req.SearchType == "" {
		return errors.WithStack(errno.ResourceNotExist)
	}
	
	posts, err := p.postDao.SearchPost(req.GetContent(), req.SearchType)
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
