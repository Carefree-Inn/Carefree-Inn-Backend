package repository

import (
	"context"
	"github.com/pkg/errors"
	"user-post/internal/repository/model"
)

type Comment struct {
	CommentId     uint32 `json:"comment_id"`
	PostId        uint32 `json:"post_id"`
	ToUserAccount string `json:"to_user_account"`
	
	CommentTime string `json:"comment_time"`
	Content     string `json:"content"`
	CommentType string `json:"comment_type"`
	
	FromUserAccount  string `json:"from_user_account"`
	FromUserAvatar   string `json:"from_user_avatar"`
	FromUserNickName string `json:"from_user_nick_name"`
}

func (c *Comment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Comment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

func (up *UserPost) MakeComment(comment *model.Comment, nickname, avatar string) error {
	session := up.db.Begin()
	if err := session.Table(comment.Table()).Create(comment).Error;
		err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	if err := session.Table(comment.Table()).Exec(
		"UPDATE post SET comments = comments + 1 WHERE post_id = ?", comment.PostId).
		Error; err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	session.Commit()
	
	repoComment := &Comment{
		CommentId:     comment.CommentId,
		ToUserAccount: comment.ToUserId,
		CommentTime:   comment.CreateTime.Format("2006-01-02 15:04:05"),
		Content:       comment.Content,
		PostId:        comment.PostId,
		
		FromUserAvatar:   comment.FromUserId,
		FromUserNickName: nickname,
		FromUserAccount:  comment.FromUserId,
		CommentType:      "make",
	}
	
	data, err := repoComment.Marshal()
	if err != nil {
		return errors.WithStack(err)
	}
	
	// comment 使用消费者-订阅者模式是为了对用户进行通知
	return errors.WithStack(up.rdb.Publish(context.TODO(), "comment_after", data).Err())
}

func (up *UserPost) DeleteComment(comment *model.Comment) error {
	session := up.db.Begin()
	if err := session.Table(comment.Table()).Delete(comment).Error;
		err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	if err := session.Table(comment.Table()).Exec(
		"UPDATE post SET comments = comments - 1 WHERE post_id = ?", comment.PostId).
		Error; err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	return errors.WithStack(session.Commit().Error)
}

func (up *UserPost) GetCommentOfPost(postId uint32, page, limit uint32) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0, limit)
	if err := up.db.Table(model.Comment{}.Table()).Where(
		"post_id=?", postId).Offset(int((page - 1) * limit)).Limit(int(limit)).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (up *UserPost) GetCommentOfUser(account string, page, limit uint32) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0, limit)
	if err := up.db.Table(model.Comment{}.Table()).Where(
		"from_id=?", account).Offset(int((page - 1) * limit)).Limit(int(limit)).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
