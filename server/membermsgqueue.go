package server

func AddMemberMsgScanFun(me *Member) {
	go handleMemberMsg(me)
}

func GetMemberMsgQueue(gid string, uid string) (chan *Msg, error) {
	member, err := GetMember(uid, gid)
	if err != nil {
		return nil, err
	}

	return member.MsgQueue, nil
}
