package server

import (
	"log"

	"github.com/gorilla/websocket"
)

type Member struct {
	Uid       string
	Name      string
	Gid       string
	Conn      *websocket.Conn
	LoginTime string
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

		oldmember.Conn.Close()

	}

	memberqueue.m[member.Uid] = member
	groupqueue.g[member.Gid] = memberqueue

	//add mermber msg queue
	AddMemberMsgQueue(member.Gid, member.Uid)
	AddMemberMsgScanFun(member.Gid, member.Uid)
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

	DeleMemberMsgQueue(member.Gid, member.Uid)
	log.Println("member : ", member.Gid+" "+member.Uid+" Resource Cleaned")
}
