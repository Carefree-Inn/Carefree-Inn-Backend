package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/go-acme/lego/v4/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (n *notificationHandler) SendNotification(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println()
		return
	}
	
	consumerLike := n.client.Subscribe(c.Request.Context(), "like")
	consumerComment := n.client.Subscribe(c.Request.Context(), "comment")
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
