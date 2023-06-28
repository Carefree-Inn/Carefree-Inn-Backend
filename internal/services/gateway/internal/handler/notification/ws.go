package notification

import (
	"context"
	"encoding/json"
	"gateway/internal"
	"gateway/pkg"
	"gateway/pkg/errno"
	"gateway/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	pb "user-post/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type LikeInfo struct {
	PostId uint32 `json:"post_id"`
	
	ToUserAccount string `json:"to_user_account"`
	ToUserAvatar  string `json:"to_user_avatar"`
	CreateTime    string `json:"create_time"`
	LikeType      string `json:"like_type"`
	
	FromUserAccount  string `json:"from_user_account"`
	FromUserAvatar   string `json:"from_user_avatar"`
	FromUserNickname string `json:"from_user_nickname"`
}

func (l *LikeInfo) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

func (l *LikeInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, l)
}

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
	IsToPost         bool   `json:"is_to_post"`
}

func (c *Comment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Comment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

//  SendNotification sendNotification
//	@Summary		获取用户通知（被点赞/评论） api
//	@Tags			notification
//	@Description	获取用户通知
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Success		200				{object}	internal.Response
//	@Router			/notification [get]
func (n *notificationHandler) SendNotification(c *gin.Context) {
	account := c.MustGet("account")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer ws.Close()
	if err != nil {
		log.Warn(nil, err)
		return
	}
	
	consumerLike := n.client.Subscribe(c.Request.Context(), "like_after")
	consumerComment := n.client.Subscribe(c.Request.Context(), "comment_after")
	defer consumerLike.Close()
	defer consumerComment.Close()
	
	for {
		select {
		case msg, ok := <-consumerLike.Channel():
			if !ok {
				continue
			}
			var likeInfo LikeInfo
			if err := likeInfo.Unmarshal([]byte(msg.Payload)); err != nil {
				log.Warn(nil, err, "Unmarshal err:")
				continue
			}
			
			if likeInfo.ToUserAccount != account {
				continue
			}
			
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Warn(nil, err, ",Like: fail to send message with websocket:")
				continue
			}
		
		case msg, ok := <-consumerComment.Channel():
			if !ok {
				continue
			}
			var comment Comment
			if err := comment.Unmarshal([]byte(msg.Payload)); err != nil {
				log.Warn(nil, err, "Unmarshal err:")
				continue
			}
			
			if comment.ToUserAccount != account {
				continue
			}
			
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Warn(nil, err, "Comment: fail to send message with websocket:")
				continue
			}
		case <-c.Request.Context().Done():
			break
		}
	}
	
}

type Notification struct {
	ActionType       string `json:"action_type"`
	FromUserAccount  string `json:"from_user_account"`
	FromUserNickname string `json:"from_user_nickname"`
	FromUserAvatar   string `json:"from_user_avatar"`
	ToUserAccount    string `json:"to_user_account"`
	PostId           uint32 `json:"post_id"`
	
	ActionId       uint32 `json:"action_id"`
	ActionTime     string `json:"action_time"`
	CommentContent string `json:"comment_content"`
	IsToPost       bool   `json:"is_to_post"`
}

//  GetNotificationHistory getNotificationHistory
//	@Summary		获取历史通知 api
//	@Tags			notification
//	@Description	获取历史通知
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"用户token"
//	@Param			page			query		int		false	"页码"
//	@Param			limit			query		int		false	"条数"
//	@Success		200				{object}	internal.Response
//	@Router			/notification/history [get]
func (n *notificationHandler) GetNotificationHistory(c *gin.Context) {
	pageStr, limitStr := c.DefaultQuery("page", "1"), c.DefaultQuery("limit", "10")
	
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		internal.Error(c, errno.ParamDataError)
		log.Warn(log.WithField("X-Request-Id", c.MustGet("uuid")), errno.ParamDataError)
		return
	}
	
	account := c.MustGet("account").(string)
	ctx := context.WithValue(c.Request.Context(), "X-Request-Id", pkg.GetUUid(c))
	resp, err := n.UserPostService.GetNotificationHistory(ctx, &pb.GetNotificationRequest{
		Account: account,
		Page:    uint32(page),
		Limit:   uint32(limit),
	})
	
	data := make([]*Notification, 0, len(resp.Notifications))
	for _, val := range resp.Notifications {
		data = append(data, &Notification{
			ActionType:       val.ActionType,
			FromUserAccount:  val.FromUserAccount,
			FromUserNickname: val.FromUserNickname,
			FromUserAvatar:   val.FromUserAvatar,
			ToUserAccount:    val.ToUserAccount,
			PostId:           val.PostId,
			ActionId:         val.ActionId,
			ActionTime:       val.ActionTime,
			CommentContent:   val.CommentContent,
			IsToPost:         val.IsToPost,
		})
	}
	
	if err != nil {
		internal.ServerError(c, errno.InternalServerError.Error())
		return
	}
	
	internal.Success(c, data)
}
