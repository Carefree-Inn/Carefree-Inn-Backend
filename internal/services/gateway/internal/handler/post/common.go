package post

import (
	"context"
	pbPost "post/proto"
	pbUser "user/proto"
)

func (p *postHandler) AssemblePostAndUser(ctx context.Context, posts []*pbPost.PostInfo) ([]*PostInfo, error) {
	var accounts = make([]string, 0, len(posts))
	for _, v := range posts {
		accounts = append(accounts, v.Account)
	}
	
	data, err := p.GetUserInfoWithAny(ctx, accounts, posts, p.assemble)
	if err != nil {
		return nil, err
	}
	return data.([]*PostInfo), nil
}

func (p *postHandler) assemble(ref map[string]*userInfo, unknown any) (any, error) {
	posts, ok := unknown.([]*pbPost.PostInfo)
	if !ok {
		return nil, nil
	}
	
	var data = make([]*PostInfo, 0, len(posts))
	for _, v := range posts {
		tag := make([]*tagInfo, 0, len(v.Tag))
		for _, val := range v.Tag {
			tag = append(tag, &tagInfo{
				TagId: val.TagId,
				Title: val.Title,
			})
		}
		
		data = append(data, &PostInfo{
			PostId:  v.PostId,
			Title:   v.Title,
			Content: v.Content,
			Category: categoryInfo{
				CategoryId: v.Category.CategoryId,
				Title:      v.Category.Title,
			},
			Tags:       tag,
			UserInfo:   ref[v.Account],
			CreateTime: v.CreateTime,
			Likes:      v.Likes,
			Liked:      v.Liked,
		})
	}
	
	return data, nil
}

func (p *postHandler) GetUserInfoWithAny(ctx context.Context, accounts []string, data any, assemble func(map[string]*userInfo, any) (any, error)) (any, error) {
	resp, err := p.UserService.GetBatchUserProfile(ctx, &pbUser.BatchUserProfileRequest{
		Accounts: accounts,
	})
	if err != nil {
		return nil, err
	}
	
	var ref = make(map[string]*userInfo, len(resp.Data))
	for _, v := range resp.Data {
		ref[v.Account] = &userInfo{
			Account:  v.Account,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Sex:      int8(v.Sex),
		}
	}
	
	return assemble(ref, data)
	
}
