package server

type RequestEntity struct {
	Uid     string
	ToUid   string
	Msg     string
	MsgType int
}

type ResponseEntity struct {
	Uid     string
	ToUid   string
	Msg     string
	MsgType int
}

func parseRequestEntity(request *RequestEntity) (*Msg, error) {

	return &Msg{request.Uid, request.ToUid, request.Msg, request.MsgType}, nil

}
