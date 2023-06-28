package repository

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"user-post/internal/repository/model"
)

type Comment struct {
	CommentId     uint32 `json:"comment_id"`
	PostId        uint32 `json:"post_id"`
	ToUserAccount string `json:"to_user_account"`
	
	CommentTime string `json:"comment_time"`
	Content     string `json:"content"`
	
	FromUserAccount  string `json:"from_user_account"`
	FromUserAvatar   string `json:"from_user_avatar"`
	FromUserNickName string `json:"from_user_nick_name"`
	
	IsToPost  bool   `json:"is_to_post"`
	PostOwner string `json:"post_owner"`
}

func (c *Comment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Comment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

func (up *UserPost) SetNotification(action any, id uint32, createTime time.Time) error {
	notification := model.Notification{}
	switch x := action.(type) {
	case *LikeInfo:
		notification.PostId = x.PostId
		notification.FromUserAccount = x.FromUserAccount
		notification.FromUserNickname = x.FromUserNickname
		notification.FromUserAvatar = x.FromUserAvatar
		notification.ToUserAccount = x.ToUserAccount
		notification.ActionType = "like"
		notification.ActionId = id
		notification.ActionTime = createTime
	case *Comment:
		notification.PostId = x.PostId
		notification.FromUserNickname = x.FromUserNickName
		notification.FromUserAccount = x.FromUserAccount
		notification.FromUserAvatar = x.FromUserAvatar
		notification.ToUserAccount = x.ToUserAccount
		notification.ActionType = "comment"
		notification.ActionId = id
		notification.ActionTime = createTime
		notification.CommentContent = x.Content
		notification.PostOwner = x.PostOwner
		notification.IsToPost = x.IsToPost
	}
	
	tx := up.db.Begin()
	if err := tx.Table("notification").Create(&notification).Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	return errors.WithStack(tx.Commit().Error)
}

func (up *UserPost) GetAccount(postId uint32) (string, error) {
	type Data struct {
		Account string `gorm:"column:account"`
	}
	data := Data{}
	if err := up.db.Table("post").Select("account").Where("post_id = ?", postId).
		First(&data).Error; err != nil {
		return "", errors.WithStack(err)
	}
	
	return data.Account, nil
}

func (up *UserPost) MakeComment(comment *model.Comment, commentInfo *Comment) error {
	session := up.db.Begin()
	if err := session.Table(comment.Table()).Create(comment).Error;
		err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	account, err := up.GetAccount(comment.PostId)
	if err != nil {
		return err
	}
	
	if err := session.Table(comment.Table()).Exec(
		"UPDATE post SET comments = comments + 1 WHERE post_id = ?", comment.PostId).
		Error; err != nil {
		return errors.WithStack(session.Rollback().Error)
	}
	
	session.Commit()
	
	commentInfo.CommentId = comment.CommentId
	commentInfo.CommentTime = comment.CreateTime.Format("2006-01-02 15:04-05")
	// isToPost 为true PostOwner才作为被通知对象
	commentInfo.IsToPost = comment.IsTop
	commentInfo.PostOwner = account
	
	if err := up.SetNotification(commentInfo, commentInfo.CommentId, comment.CreateTime); err != nil {
		return err
	}
	
	data, err := commentInfo.Marshal()
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
		Order("create_time desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (up *UserPost) GetCommentOfUser(account string, page, limit uint32) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0, limit)
	if err := up.db.Table(model.Comment{}.Table()).Where(
		"from_id=?", account).Offset(int((page - 1) * limit)).Limit(int(limit)).
		Order("create_time desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
