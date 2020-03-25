package server

import (
	"log"
	"sync"
)

type Msg struct {
	Uid     string
	ToUid   string
	Msg     string
	MsgType int
}

//var MemberQueue map[string]*Member

type MemberQueue struct {
	sync.RWMutex
	m map[string]*Member
}

var memberqueue MemberQueue

var MsgQueue chan *Msg

var MsgQueusMAX = 20

func init() {
	memberqueue.m = make(map[string]*Member)
	//MemberQueue = make(map[string]*Member)
	MsgQueue = make(chan *Msg, MsgQueusMAX)
	go HandleMsg()
}

func AddMember(member *Member) {
	memberqueue.RLock()
	defer memberqueue.RUnlock()

	memberqueue.m[member.Uid] = member
}

func GetAllOnLineMember() map[string]*Member {
	return memberqueue.m
}

func GetMember(uid string) *Member {
	// memberqueue.RLock()
	// defer memberqueue.RUnlock()

	if _, ok := memberqueue.m[uid]; ok {
		return memberqueue.m[uid]
	}

	return nil
}

func DeleMember(uid string) {
	memberqueue.RLock()
	defer memberqueue.RUnlock()
	delete(memberqueue.m, uid)
}

func AddMsg(msg *Msg) {
	MsgQueue <- msg
}

func HandleMsg() {
	for {
		msg := <-MsgQueue
		go sendMsg(msg)
	}
}

func sendMsg(msg *Msg) {
	switch msg.MsgType {
	case P2P:
		sendP2PMsg(msg)
	case P2A:
		sendP2AMsg(msg)
	case N2A:
		sendN2AMsg(msg)
	case N2P:
		sendN2PMsg(msg)
	default:
		return
	}
}

func sendP2PMsg(msg *Msg) {
	rm := GetMember(msg.ToUid)
	if rm == nil {
		if msg.Uid != msg.ToUid {
			log.Println("member : ", msg.ToUid, " not on line or no existence")
			msg.ToUid = msg.Uid
			msg.Msg = "member : " + msg.ToUid + " not on line or no existence"
			sendN2PMsg(msg)
		}
		return
	}
	con := rm.Conn
	res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Msg, msg.MsgType}
	err := con.WriteJSON(res)
	if err != nil {
		log.Printf("write fail = %v\n", err)
		return
	}
}

func sendP2AMsg(msg *Msg) {
	msg.ToUid = "ALL"
	members := GetAllOnLineMember()
	for _, v := range members {
		go func(rm *Member) {
			con := rm.Conn
			res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Msg, msg.MsgType}
			err := con.WriteJSON(res)
			if err != nil {
				log.Printf("write fail = %v\n", err)
				return
			}
		}(v)
	}
}
func sendN2AMsg(msg *Msg) {
	msg.Uid = "NOTICE"
	msg.ToUid = "ALL"
	members := GetAllOnLineMember()
	for _, v := range members {
		go func(rm *Member) {
			con := rm.Conn
			res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Msg, msg.MsgType}
			err := con.WriteJSON(res)
			if err != nil {
				log.Printf("write fail = %v\n", err)
				return
			}
		}(v)
	}
}
func sendN2PMsg(msg *Msg) {
	msg.ToUid = msg.Uid
	msg.Uid = "NOTICE"
	sendP2PMsg(msg)
}
