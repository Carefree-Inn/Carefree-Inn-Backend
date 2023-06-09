package handler

import (
	"context"
	"user-post/internal/repository"
	"user-post/internal/repository/model"
	pb "user-post/proto"
)

func (up *UserPostService) MakeComment(ctx context.Context, req *pb.MakeCommentRequest, resp *pb.Response) error {
	return up.userPostDao.MakeComment(
		&model.Comment{
			PostId:       req.PostId,
			FromUserId:   req.FromUserAccount,
			ToUserId:     req.ToUserAccount,
			Content:      req.Content,
			IsTop:        req.IsTop,
			TopCommentId: req.TopCommentId,
		}, &repository.Comment{
			PostId:           req.PostId,
			ToUserAccount:    req.ToUserAccount,
			Content:          req.Content,
			FromUserAccount:  req.FromUserAccount,
			FromUserAvatar:   req.FromUserAvatar,
			FromUserNickName: req.FromUserNickName,
		})
}

func (up *UserPostService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest, resp *pb.Response) error {
	return up.userPostDao.DeleteComment(&model.Comment{
		CommentId: req.CommentId,
	})
}

func (up *UserPostService) GetCommentOfPost(ctx context.Context, req *pb.GetCommentOfPostRequest, resp *pb.CommentOfPostResponse) error {
	comments, err := up.userPostDao.GetCommentOfPost(req.PostId, req.Page, req.Limit)
	if err != nil {
		return err
	}
	resp.PostId = req.PostId
	topComments := make([]*pb.TopCommentResponse, 0, req.Limit)
	index := make(map[uint32]int, 0)
	start := 0
	
	for _, v := range comments {
		if v.IsTop {
			topComments = append(topComments, &pb.TopCommentResponse{
				CommentId:  v.CommentId,
				FromUserId: v.FromUserId,
				Content:    v.Content,
				CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
				Comments:   make([]*pb.CommentResponse, 0),
			})
			index[v.CommentId] = start
			start++
		}
	}
	
	for _, v := range comments {
		if !v.IsTop {
			// 顺序错误，一开始就是没有top的怎么办
			topComments[index[v.TopCommentId]].Comments = append(topComments[index[v.TopCommentId]].Comments,
				&pb.CommentResponse{
					CommentId:  v.CommentId,
					FromUserId: v.FromUserId,
					ToUserId:   v.ToUserId,
					Content:    v.Content,
					CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
				})
		}
	}
	
	resp.Comments = topComments
	return nil
}

func (up *UserPostService) GetCommentOfUser(ctx context.Context, req *pb.GetCommentOfUserRequest, resp *pb.CommentOfUserResponse) error {
	comments, err := up.userPostDao.GetCommentOfUser(req.Account, uint32(req.Page), uint32(req.Limit))
	if err != nil {
		return err
	}
	resp.Account = req.Account
	data := make([]*pb.UserComment, 0, len(comments))
	for _, one := range comments {
		comment := &pb.UserComment{
			CommentId:  one.CommentId,
			FromUserId: one.FromUserId,
			ToUserId:   one.ToUserId,
			Content:    one.Content,
			CreateTime: one.CreateTime.Format("2006-01-02 15:04:05"),
			
			PostId:       one.PostId,
			IsTop:        one.IsTop,
			TopCommentId: one.TopCommentId,
		}
		data = append(data, comment)
	}
	resp.Comments = data
	
	return nil
}
