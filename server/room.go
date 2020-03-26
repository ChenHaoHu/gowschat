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

	MsgQueue = make(chan *Msg, MsgQueusMAX)
	go HandleMsg()
}

func AddMember(member *Member, gid string) {

	memberqueue := groupqueue.g[gid]

	if memberqueue == nil {
		memberqueue = &MemberQueue{m: make(map[string]*Member)}
	}

	memberqueue.RLock()
	defer memberqueue.RUnlock()

	oldmember := memberqueue.m[member.Uid]
	//	log.Println(oldmember)

	if oldmember != nil {
		//close conn

		// sendN2PMsg(&Msg{Uid: oldmember.Uid, ToUid: oldmember.Uid, Gid: gid,
		// 	Msg: "Your account is repeatedly logged in in this group", SendType: N2P})

		oldmember.Conn.Close()

	}

	memberqueue.m[member.Uid] = member
	groupqueue.g[gid] = memberqueue

}

func GetAllOnLineMember() map[string]*MemberQueue {
	return groupqueue.g
}

func GetGropOnLineMember(gid string) map[string]*Member {
	return groupqueue.g[gid].m
}

func GetMember(uid string, gid string) *Member {
	// memberqueue.RLock()
	// defer memberqueue.RUnlock()

	memberqueue := groupqueue.g[gid]

	if memberqueue == nil {
		return nil
	}

	if _, ok := memberqueue.m[uid]; ok {
		return memberqueue.m[uid]
	}

	return nil
}

func DeleMember(uid string, gid string) {

	memberqueue := groupqueue.g[gid]

	if memberqueue == nil {
		return
	}

	memberqueue.RLock()
	defer memberqueue.RUnlock()
	delete(memberqueue.m, uid)

	if len(memberqueue.m) == 0 {
		delete(groupqueue.g, gid)
	}
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

func sendP2PMsg(msg *Msg) {
	rm := GetMember(msg.ToUid, msg.Gid)
	if rm == nil {
		if msg.Uid != msg.ToUid {
			log.Println("member : ", msg.ToUid, " in ", msg.Gid, " not on line or no existence")
			msg.ToUid = msg.Uid
			msg.Msg = "member : " + msg.ToUid + " in " + msg.Gid + " not on line or no existence"
			sendN2PMsg(msg)
		}
		return
	}
	con := rm.Conn
	res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Gid, msg.Msg, msg.SendType}
	err := con.WriteJSON(res)
	if err != nil {
		log.Printf("write fail = %v\n", err)
		return
	}
}

func sendP2GMsg(msg *Msg) {
	msg.ToUid = "GroupALL"
	members := GetGropOnLineMember(msg.Gid)
	for _, v := range members {
		go func(rm *Member) {
			con := rm.Conn
			res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Gid, msg.Msg, msg.SendType}
			err := con.WriteJSON(res)
			if err != nil {
				log.Printf("write fail = %v\n", err)
				return
			}
		}(v)
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
			go func(rm *Member) {
				con := rm.Conn
				res := &ResponseEntity{msg.Uid, msg.ToUid, msg.Gid, msg.Msg, msg.SendType}
				err := con.WriteJSON(res)
				if err != nil {
					log.Printf("write fail = %v\n", err)
					return
				}
			}(i)
		}
	}

	//defer recover()
}

func sendN2PMsg(msg *Msg) {
	//msg.ToUid = msg.Uid
	msg.Uid = "NOTICE"
	sendP2PMsg(msg)
}
