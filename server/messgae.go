package server

const MsgQueueMax = 100

type Msg struct {
	Uid      string
	ToUid    string
	Gid      string
	Msg      string
	SendType int
}

type MsgQueue struct {
	Queue   chan Msg
	Total   int
	IsClose bool
}

func NewMsgQueue() (*MsgQueue, error) {
	queue := make(chan Msg, MsgQueueMax)
	msgQueue := &MsgQueue{
		Queue: queue,
		Total: 0,
	}
	return msgQueue, nil
}

func (msgQueue *MsgQueue) AddMsg(msg *Msg) error {

	msgQueue.Queue <- *msg
	return nil
}
