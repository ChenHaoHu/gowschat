package server

type RequestEntity struct {
	ToUid    string
	Msg      string
	SendType int
}

type ResponseEntity struct {
	Uid      string
	ToUid    string
	Gid      string
	Msg      string
	SendType int
}

func parseRequestEntity(request *RequestEntity, member *Member, gid string) (*Msg, error) {

	return &Msg{
		Uid:      member.Uid,
		ToUid:    request.ToUid,
		Gid:      gid,
		Msg:      request.Msg,
		SendType: request.SendType,
	}, nil

}
