package server

import (
	"log"
	"sync"
)

type Msg struct {
	Uid   string
	ToUid string
	Gid   string
	Msg   string

	SendType int
}

//var MemberQueue map[string]*Member

type GroupQueue struct {
	sync.RWMutex
	g map[string]*MemberQueue
}

type MemberQueue struct {
	sync.RWMutex
	m map[string]*Member
}

//var memberqueue MemberQueue
var groupqueue GroupQueue

var MsgQueue chan *Msg

var MsgQueusMAX = 20

func init() {
	//memberqueue.m = make(map[string]*Member)

	//MemberQueue = make(map[string]*Member)
	groupqueue.g = make(map[string]*MemberQueue)
	groupqueue.g["ALL"] = &MemberQueue{m: make(map[string]*Member)}

	groupmsgqueue.g = make(map[string]*MemberMsgQueue)

	MsgQueue = make(chan *Msg, MsgQueusMAX)
	go HandleMsg()
}

func AddMsg(msg *Msg) {
	MsgQueue <- msg
}

func HandleMsg() {
	for {
		msg := <-MsgQueue
		//consider need or not need be go
		sendMsg(msg)
	}
}

func sendMsg(msg *Msg) {

	switch msg.SendType {
	case P2P:
		sendP2PMsg(msg)
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

func handleMemberMsg(gid string, uid string) {
	membermsgqueue, _ := GetMemberMsgQueue(gid, uid)
	//get member  conn
	member := GetMember(uid, gid)

	for {
		msg, flag := <-membermsgqueue
		if flag == false {
			log.Println("close membermsgqueue :", gid, "  ", uid)
			return
		}
		log.Println(gid, uid, " 收到来自 ", msg.Uid, " 的消息")

		sendP2PMsgDirect(msg, member)
	}
}

func sendP2PMsgDirect(msg *Msg, member *Member) {
	con := member.Conn
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
		log.Println(res.Uid, " 成功发送消息给 ", res.ToUid, " : "+res.Msg)
	}
}

func sendP2PMsg(msg *Msg) {
	// member := GetMember(msg.ToUid, msg.Gid)
	// if member == nil {
	// 	if msg.Uid != msg.ToUid {
	// 		log.Println("member : ", msg.ToUid, " in ", msg.Gid, " not on line or no existence")
	// 		msg.ToUid = msg.Uid
	// 		msg.Msg = "member : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence"
	// 		sendN2PMsg(msg)
	// 	}
	// 	return
	// }
	//get member msg channel
	membermsgqueue, err := GetMemberMsgQueue(msg.Gid, msg.ToUid)
	if err != nil {
		//send notice
		msg.Msg = "member : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence"
		msg.ToUid = msg.Uid
		msg.SendType = N2P
		AddMsg(msg)
		return
	}
	membermsgqueue <- msg
}

func sendP2GMsg(msg *Msg) {
	msg.ToUid = "GroupALL"
	members := GetGropOnLineMember(msg.Gid)
	for _, v := range members {
		//go sendP2PMsgDirect(msg, v)
		membermsgqueue, err := GetMemberMsgQueue(v.Gid, v.Uid)
		if err != nil {
			//send notice
			msg.Msg = "member : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence"
			if msg.ToUid == msg.Uid {
				return
			}
			msg.ToUid = msg.Uid
			msg.SendType = N2P
			if msg.ToUid == msg.Uid {
				return
			}
			AddMsg(msg)
			return
		}
		membermsgqueue <- msg
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
	groups := GetAllOnLineMember()
	for _, v := range groups {
		// log.Println(len(v.m))
		for _, i := range v.m {
			membermsgqueue, err := GetMemberMsgQueue(i.Gid, i.Uid)
			if err != nil {
				//send notice
				msg.Msg = "member : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence"
				msg.ToUid = msg.Uid
				msg.SendType = N2P
				if msg.ToUid == msg.Uid {
					return
				}
				AddMsg(msg)
				return
			}
			membermsgqueue <- msg
		}
	}

	//defer recover()
}

func sendN2PMsg(msg *Msg) {
	//msg.ToUid = msg.Uid
	msg.Uid = "NOTICE"
	sendP2PMsg(msg)
}
