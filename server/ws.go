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
	member, gid, err := ParseToken(userIdent.Token)
	if err != nil {
		log.Println(err)
		return
	}
	member.Conn = conn
	AddMember(member, gid)
	log.Println("member name :", member.Name, " member gid :", gid, " member uid :", member.Uid, " login ok")

	//check EnableNotifyAllWhenLogin
	if EnableNotifyAllWhenLogin {
		msg := &Msg{member.Uid, member.Uid, gid, " member name :" + member.Name + " member uid :" +
			member.Uid + " login ok", N2G}
		AddMsg(msg)
	}

	defer func(member *Member) {
		if EnableNotifyAllWhenLogout {
			msg := &Msg{member.Uid, member.Uid, gid, "member name :" + member.Name + " member uid :" +
				member.Uid + " logout ok", N2G}
			AddMsg(msg)
		}
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "json err"))
		conn.Close()

		defer recover()

	}(member)

	for {
		d := &RequestEntity{}
		err := conn.ReadJSON(d)

		if err != nil {
			DeleMember(member.Uid, gid)
			log.Println("member name :", member.Name, " member uid :", member.Uid, " logout ok")

			log.Printf("read fail = %v\n", err)
			return

			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {

			// } else {
			// 	// log.Printf("read fail = %v\n", err)
			// 	// return
			// 	msg := &Msg{member.Uid, member.Uid, gid, "Request Entity is not json or wrong format", N2A}
			// 	AddMsg(msg)
			// 	continue
			// }

		}

		//check authority
		err = CheckInden(d, member, gid)

		var msg *Msg

		if err != nil {
			msg = &Msg{member.Uid, member.Uid, gid, err.Error(), N2P}
		} else {
			msg, err = parseRequestEntity(d, member, gid)
			if err != nil {
				log.Println(err)
				return
			}

		}

		AddMsg(msg)
		// if d.Msg != "" {
		// 	//log.Println("add msg")
		// 	MsgQueue <- msg
		// }

	}

}
