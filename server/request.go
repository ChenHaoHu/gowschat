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

func parseRequestEntity(request *RequestEntity, Client *Client) (*Msg, error) {

	return &Msg{
		Uid:      Client.Uid,
		ToUid:    request.ToUid,
		Gid:      Client.Gid,
		Msg:      request.Msg,
		SendType: request.SendType,
	}, nil

}
