package server

import (
	"log"
	"sync"
)

type Msg struct {
	Uid   string
	ToUid string
	Msg   string
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
	go HanleMsg()
}

func AddMember(member *Member) {
	memberqueue.RLock()
	defer memberqueue.RUnlock()

	memberqueue.m[member.Uid] = member
}

func GetMember(uid string) *Member {
	memberqueue.RLock()
	defer memberqueue.RUnlock()

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

func HanleMsg() {
	for {
		msg := <-MsgQueue
		go sendMsg(msg)
	}
}

func sendMsg(msg *Msg) {
	rm := GetMember(msg.ToUid)
	if rm == nil {
		if msg.Uid != msg.ToUid {
			log.Println("member : ", msg.ToUid, " not on line or no existence")
			msg.ToUid = msg.Uid
			msg.Msg = "member : " + msg.ToUid + " not on line or no existence"
			sendMsg(msg)
		}
		return
	}
	con := rm.Conn
	res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Msg}
	err := con.WriteJSON(res)
	if err != nil {
		log.Printf("write fail = %v\n", err)
		return
	}
}
