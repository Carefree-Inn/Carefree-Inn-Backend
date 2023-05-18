package handler

import (
	"context"
	"user-post/internal/repository/model"
	pb "user-post/proto"
)

func (up *UserPostService) MakeComment(ctx context.Context, req *pb.MakeCommentRequest, resp *pb.Response) error {
	return up.userPostDao.MakeComment(
		&model.Comment{
			IsTop:        req.IsTop,
			TopCommentId: req.TopCommentId,
			FromUserId:   req.FromUserId,
			ToUserId:     req.ToUserId,
			Content:      req.Content,
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
		} else {
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
