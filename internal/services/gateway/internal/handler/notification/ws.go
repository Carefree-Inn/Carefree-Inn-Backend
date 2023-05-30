package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/go-acme/lego/v4/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

//  SendNotification sendNotification
//	@Summary		获取用户通知（被点赞/评论） api
//	@Tags			post
//	@Description	获取用户通知
//	@Accept			json
//	@Produce		json
//	@Param			Authorzation	header		string	true	"用户token"
//	@Success		200				{object}	internal.Response
//	@Router			/notification [get]
func (n *notificationHandler) SendNotification(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Like: fail to send message with websocket:", err)
			}
			log.Println(msg.Payload)
		case msg, ok := <-consumerComment.Channel():
			if !ok {
				continue
			}
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Comment: fail to send message with websocket:", err)
			}
			log.Println(msg.Payload)
		case <-c.Request.Context().Done():
			return
		}
	}
	
	defer ws.Close()
	
}
