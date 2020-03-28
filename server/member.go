package server

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

type Member struct {
	Uid       string
	Name      string
	Gid       string
	Conn      *websocket.Conn
	LoginTime string
	MsgQueue  chan *Msg
}

func AddMember(member *Member) {

	memberqueue := groupqueue.g[member.Gid]

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
		DeleMember(oldmember)
		// oldmember.Conn.Close()

	}

	// init member msg
	member.MsgQueue = make(chan *Msg, MsgQueusMAX)

	memberqueue.m[member.Uid] = member
	groupqueue.g[member.Gid] = memberqueue

	//add mermber msg queue

	// member add msg field
	//AddMemberMsgQueue(member.Gid, member.Uid)
	AddMemberMsgScanFun(member)
}

func GetAllOnLineMember() map[string]*MemberQueue {
	return groupqueue.g
}

func GetGropOnLineMember(gid string) map[string]*Member {

	if groupqueue.g[gid] == nil {
		return nil
	}

	return groupqueue.g[gid].m
}

func GetMember(uid string, gid string) (*Member, error) {
	// memberqueue.RLock()
	// defer memberqueue.RUnlock()

	memberqueue := groupqueue.g[gid]

	if memberqueue == nil {
		return nil, errors.New("no group")
	}

	if _, ok := memberqueue.m[uid]; ok {
		return memberqueue.m[uid], nil
	}

	return nil, errors.New("no member")
}

func DeleMember(member *Member) {
	gid := member.Gid
	uid := member.Uid
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

	//DeleMemberMsgQueue(member.Gid, member.Uid)

	close(member.MsgQueue)

	log.Println("member : ", member.Gid+" "+member.Uid+" Resource Cleaned")
}

func (member *Member) SendMsg(msg *Msg) {
	membermsgqueue := member.MsgQueue
	membermsgqueue <- msg
}
