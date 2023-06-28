package comment

import (
	"context"
	"gateway/internal/service"
	pb "user-post/proto"
	pbUser "user/proto"
)

type commentHandler struct {
	pb.UserPostService
	pbUser.UserService
}

func NewCommentHandler() *commentHandler {
	return &commentHandler{
		service.NewUserPostService(),
		service.NewUserService(),
	}
}

func (c *commentHandler) Convert(user *pbUser.UserProfileResponse) *UserInfo {
	return &UserInfo{
		Account:  user.Account,
		Nickname: user.Nickname,
		Sex:      int8(user.Sex),
		Avatar:   user.Avatar,
	}
}

func (c *commentHandler) UpperAssemble(ctx context.Context, response *pb.CommentOfPostResponse) ([]*Comment, error) {
	ids := make([]string, 0)
	for _, v := range response.Comments {
		ids = append(ids, v.FromUserId)
		for _, val := range v.Comments {
			ids = append(ids, val.FromUserId, val.ToUserId)
		}
	}
	
	return c.Assemble(ctx, ids, response)
}

func (c *commentHandler) Assemble(ctx context.Context, userIds []string, response *pb.CommentOfPostResponse) ([]*Comment, error) {
	resp, err := c.UserService.GetBatchUserProfile(ctx, &pbUser.BatchUserProfileRequest{
		Accounts: userIds,
	})
	if err != nil {
		return nil, err
	}
	
	var comments = make([]*Comment, 0, len(resp.Data))
	for _, val := range response.Comments {
		comment := &Comment{
			CommentId:       val.CommentId,
			PostId:          response.PostId,
			Content:         val.Content,
			CreateTime:      val.CreateTime,
			FromUserAccount: c.Convert(resp.Data[val.FromUserId]),
			ToUserAccount:   nil,
			IsTop:           true,
			TopCommentId:    0,
			ChildComment:    make([]*Comment, 0),
		}
		for _, v := range val.Comments {
			comment.ChildComment = append(comment.ChildComment,
				&Comment{
					CommentId:       v.CommentId,
					PostId:          response.PostId,
					Content:         v.Content,
					CreateTime:      v.CreateTime,
					FromUserAccount: c.Convert(resp.Data[v.FromUserId]),
					ToUserAccount:   c.Convert(resp.Data[v.ToUserId]),
					IsTop:           false,
					TopCommentId:    val.CommentId,
					ChildComment:    nil,
				})
		}
		comments = append(comments, comment)
	}
	
	return comments, nil
}
