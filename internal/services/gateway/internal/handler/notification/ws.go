package notification

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type LikeInfo struct {
	PostId uint32 `json:"post_id"`
	
	ToUserAccount   string `json:"to_user_account"`
	CreateTime      string `json:"create_time"`
	LikeType        string `json:"like_type"`
	FromUserAccount string `json:"from_user_account"`
	FromUserAvatar  string `json:"from_user_avatar"`
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
		log.Println(err)
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
				log.Println("Unmarshal err:", err)
				continue
			}
			
			if likeInfo.ToUserAccount != account {
				continue
			}
			
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Like: fail to send message with websocket:", err)
				continue
			}
		
		case msg, ok := <-consumerComment.Channel():
			if !ok {
				continue
			}
			var comment Comment
			if err := comment.Unmarshal([]byte(msg.Payload)); err != nil {
				log.Println("Unmarshal err:", err)
				continue
			}
			
			if comment.ToUserAccount != account {
				continue
			}
			
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Comment: fail to send message with websocket:", err)
				continue
			}
		case <-c.Request.Context().Done():
			break
		}
	}
	
}
