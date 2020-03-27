package server

import (
	"errors"
	"sync"
)

type GroupMsgQueue struct {
	sync.RWMutex
	g map[string]*MemberMsgQueue
}

type MemberMsgQueue struct {
	sync.RWMutex
	m map[string]chan *Msg
}

var groupmsgqueue GroupMsgQueue

func AddMemberMsgScanFun(gid string, uid string) {
	go handleMemberMsg(gid, uid)
}

func AddMemberMsgQueue(gid string, uid string) {
	membermsgqueue := groupmsgqueue.g[gid]

	if membermsgqueue == nil {
		membermsgqueue = &MemberMsgQueue{m: make(map[string]chan *Msg)}
	}

	msg := membermsgqueue.m[uid]

	if msg != nil {
		close(msg)
	}
	membermsgqueue.m[uid] = make(chan *Msg, MsgQueusMAX)
	groupmsgqueue.g[gid] = membermsgqueue
}

func DeleMemberMsgQueue(gid string, uid string) {
	membermsgqueue := groupmsgqueue.g[gid]

	if membermsgqueue == nil {
		return
	}

	msg := membermsgqueue.m[uid]

	if msg != nil {
		close(msg)
	}

	delete(membermsgqueue.m, uid)

	if len(membermsgqueue.m) == 0 {
		delete(groupmsgqueue.g, gid)
	}

}

func GetMemberMsgQueue(gid string, uid string) (chan *Msg, error) {
	membermsgqueue := groupmsgqueue.g[gid]

	if membermsgqueue == nil {
		return nil, errors.New("no group ")
	}

	msgchan := membermsgqueue.m[uid]

	if msgchan == nil {
		return nil, errors.New("no group ")
	}

	return msgchan, nil
}
