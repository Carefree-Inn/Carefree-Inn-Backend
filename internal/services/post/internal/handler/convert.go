package handler

import (
	"post/internal/repository/model"
	pb "post/proto"
)

func (p *PostService) convertPost(posts ...*model.Post) ([]*pb.PostInfo, error) {
	var data = make([]*pb.PostInfo, 0, len(posts))
	var t = make(map[uint32]bool)
	var ids = make([]uint32, 0, len(data))
	
	for _, v := range posts {
		t[v.CategoryId] = true
	}
	for k := range t {
		ids = append(ids, k)
	}
	
	categories, err := p.postDao.GetCategory(ids)
	if err != nil {
		return nil, err
	}
	
	var m = make(map[uint32]*model.Category)
	for _, v := range categories {
		m[v.CategoryId] = v
	}
	
	for _, v := range posts {
		var tags = make([]*pb.TagInfo, 0, len(v.Tags))
		for _, one := range v.Tags {
			tags = append(tags, &pb.TagInfo{
				Title: one.Title,
				TagId: one.TagId,
			})
		}
		
		data = append(data, &pb.PostInfo{
			Account: v.Account,
			PostId:  v.PostId,
			Title:   v.Title,
			Content: v.Content,
			Category: &pb.CategoryInfo{
				CategoryId: m[v.CategoryId].CategoryId,
				Title:      m[v.CategoryId].Title,
			},
			Likes:      v.Likes,
			Comments:   v.Comments,
			CreateTime: v.CreateTime.Format("2006-01-02 15:01:05"),
			Tag:        tags,
			Liked:      v.Liked,
		})
	}
	
	return data, nil
}
