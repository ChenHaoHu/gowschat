package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserIdent struct {
	Token string `uri:"token" binding:"required"`
}

func HandleWS(c *gin.Context) {
	r := c.Request
	w := c.Writer
	wsconn, err := upgrader.Upgrade(w, r, nil)
	conn := NewConn(wsconn)
	if err != nil {
		log.Printf("err = %s\n", err)
		return
	}

	defer func() {
		//结束消息
		//conn.Link.WriteMessage(websocket.CloseMessage,
		//	websocket.FormatCloseMessage(websocket.CloseNormalClosure, "json err"))
		conn.Close()
	}()

	//conn ok
	//get token
	var userIdent UserIdent
	if err := c.ShouldBindUri(&userIdent); err != nil {
		//c.JSON(400, gin.H{"msg": err})
		conn.Link.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	//parse token
	client, err := ParseToken(userIdent.Token)
	if err != nil {
		conn.Link.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	client.Conn = conn

	//new a msg queue
	msgQueue, _ := NewMsgQueue()
	client.MsgQueue = msgQueue

	//check EnableNotifyAllWhenLogin
	if EnableNotifyAllWhenLogin {
		msg := &Msg{client.Uid, client.Uid, client.Gid, " client name :" + client.Name + " client uid :" +
			client.Uid + " login ok", N2G}
		client.SendMsg(msg)
	}

	defer func(client *Client) {
		if EnableNotifyAllWhenLogout {
			msg := &Msg{client.Uid, client.Uid, client.Gid, "client name :" + client.Name + " client uid :" +
				client.Uid + " logout ok", N2G}

			client.SendMsg(msg)
		}

	}(client)

	group.RegisterClientToGroup(client)

	//AddClient(client)
	go client.ReciveData()
	go client.WriteData()

	<-client.CloseRequest

}
