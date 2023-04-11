package handler

import (
	"context"
	pb "post/proto"
)

func (p *PostService) GetPostOfTag(ctx context.Context, in *pb.TagInfo, resp *pb.PostResponse) error {
	posts, err := p.postDao.GetPostOfTag(in.Title)
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
