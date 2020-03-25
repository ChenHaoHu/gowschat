package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type UserIdent struct {
	Token string `uri:"token" binding:"required"`
}

func HandleWS(c *gin.Context) {
	r := c.Request
	w := c.Writer
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("err = %s\n", err)
		return
	}

	//conn ok
	//get token
	var userIdent UserIdent
	if err := c.ShouldBindUri(&userIdent); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	//parse token
	member, err := ParseToken(userIdent.Token)
	if err != nil {
		log.Println(err)
		return
	}
	member.Conn = conn
	AddMember(member)
	log.Println("member name :", member.Name, "member uid :", member.Uid, " login ok")

	defer func(member *Member) {
		DeleMember(member.Uid)
		log.Println("member name :", member.Name, "member uid :", member.Uid, " logout ok")
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "json err"))
		conn.Close()
	}(member)

	for {
		d := &RequestEntity{}
		err := conn.ReadJSON(d)

		if err != nil {
			log.Printf("read fail = %v\n", err)
			return
		}

		if d.Msg != "" {
			//log.Println("add msg")
			MsgQueue <- &Msg{d.Uid, d.ToUid, d.Msg}
		}

	}

}
