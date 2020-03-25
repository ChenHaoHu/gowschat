package server

type RequestEntity struct {
	Uid   string
	ToUid string
	Msg   string
}

type ResponseEntity struct {
	Uid   string
	ToUid string
	Msg   string
}
