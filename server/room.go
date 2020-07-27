package server

import (
	"log"
)

var group Group

func init() {
	group.g = make(map[string]*Clients)
	group.g["ALL"] = &Clients{m: make(map[string]*Client)}
}

func SendMsg(msg *Msg, client *Client) {
	log.Printf("--------------  %+v\n", msg)

	switch msg.SendType {
	case P2P:
		sendP2PMsg(msg, client)
	case P2G:
		sendP2GMsg(msg)
	case N2G:
		sendN2GMsg(msg)
	case N2A:
		sendN2AMsg(msg)
	case N2P:
		sendN2PMsg(msg)
	default:
		return
	}
}

func sendP2PMsgDirect(msg *Msg, client *Client) {
	con := client.Conn
	res := &ResponseEntity{
		Uid:      msg.Uid,
		ToUid:    msg.ToUid,
		Gid:      msg.Gid,
		Msg:      msg.Msg,
		SendType: msg.SendType,
	}
	err := con.WriteJSON(res)
	if err != nil {
		log.Printf("write fail = %v\n", err)
		return
	} else {
		log.Println("gid: ", res.Gid, "uid: ", res.Uid, " 成功发送消息给 ", res.ToUid, " : "+res.Msg)
	}
}

func sendP2PMsg(msg *Msg, client *Client) {

	sendP2PMsgDirect(msg, client)

	//if err != nil {
	//	//send notice
	//	if msg.SendType == P2P {
	//		msg.Msg = "Client : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence " + err.Error()
	//		msg.ToUid = msg.Uid
	//		msg.SendType = N2P
	//		sendN2PMsg(msg)
	//	}
	//	return
	//}

}

func sendP2GMsg(msg *Msg) {
	msg.ToUid = "GroupALL"
	clients := GetGroupOnLineClient(msg.Gid)
	for _, v := range clients {
		go sendP2PMsgDirect(msg, v)
	}
}

func sendN2GMsg(msg *Msg) {
	msg.Uid = "NOTICE"
	msg.ToUid = "GroupALL"
	sendP2GMsg(msg)
}

func sendN2AMsg(msg *Msg) {
	msg.Uid = "NOTICE"
	msg.ToUid = "ALL"
	groups := GetAllOnLineClient()
	for _, v := range groups {
		for _, i := range v.m {
			i.SendMsg(msg)
		}
	}

	//defer recover()
}

func sendN2PMsg(msg *Msg) {
	msg.Uid = "NOTICE"
	sendP2PMsg(msg, nil)
}
