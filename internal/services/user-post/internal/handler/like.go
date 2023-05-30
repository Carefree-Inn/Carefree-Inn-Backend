package handler

import (
	"context"
	pb "user-post/proto"
)

func (up *UserPostService) MakeLike(ctx context.Context, req *pb.MakeLikeRequest, resp *pb.Response) error {
	return up.userPostDao.MakeLike(req.PostId, req.Account, req.Title, req.Avatar)
}

func (up *UserPostService) CancelLike(ctx context.Context, req *pb.CancelLikeRequest, resp *pb.Response) error {
	return up.userPostDao.DeleteLike(req.PostId, req.Account)
}

//func (up *UserPostService) GetLikeOfUser(ctx context.Context, req *pb.LikeOfUserRequest, resp *pb.LikeOfUserResponse) error {
//	likes, err := up.userPostDao.GetLikes(req.Account, req.Page, req.Limit)
//	if err != nil {
//		return err
//	}
//
//	postIds := make([]uint32, 0, len(likes))
//	for _, v := range likes {
//		postIds = append(postIds, v.PostId)
//	}
//
//	resp.PostIds = postIds
//	resp.Account = req.Account
//	return nil
//}

//func (up *UserPostService) IsBatchLiked(ctx context.Context, req *pb.BactchLikedRequest, resp *pb.BatchLikedResponse) error {
//	data, err := up.userPostDao.IsBatchLiked(req.Account, req.PostIds)
//	if err != nil {
//		return err
//	}
//
//	status := make([]*pb.LikeStatus, 0, len(data))
//	for _, id := range req.PostIds {
//		in := false
//		for _, v := range data {
//			if id == v.PostId {
//				status = append(status, &pb.LikeStatus{
//					Liked:  true,
//					PostId: id,
//				})
//				in = true
//				break
//			}
//		}
//		if !in {
//			status = append(status, &pb.LikeStatus{
//				Liked:  false,
//				PostId: id,
//			})
//		}
//	}
//
//	resp.Status = status
//	resp.Account = req.Account
//	return nil
//}
