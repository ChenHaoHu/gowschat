package server

import (
	"fmt"
	"log"
)

type Client struct {
	Uid          string
	Name         string
	Gid          string
	Conn         *Conn
	LoginTime    string
	MsgQueue     *MsgQueue
	CloseWrite   chan int
	CloseRequest chan int
	IsClose      bool
}

func (client *Client) SendMsg(msg *Msg) {
	if !client.IsClose {
		client.MsgQueue.AddMsg(msg)
	}
}

func (client *Client) ReciveData() {
	conn := client.Conn
	for {
		d := &RequestEntity{}
		err := conn.ReadJSON(d)
		if err != nil {
			client.DeleteSelf()
			log.Printf("read fail = %v\n", err)
			return
		}

		//check authority
		err = CheckInden(d, client)
		var msg *Msg
		if err != nil {
			msg = &Msg{client.Uid, client.Uid, client.Gid, err.Error(), N2P}
		} else {
			msg, err = parseRequestEntity(d, client)
			if err != nil {
				log.Println(err)
				return
			}
		}
		client.SendMsg(msg)
	}
}

func (client *Client) WriteData() {
	queue := client.MsgQueue.Queue
	gid := client.Gid
	uid := client.Uid
	for {
		select {
		case msg := <-queue:
			log.Println("gid: ", gid, "uid: ", uid, " 收到来自 ", msg.Uid, " 的消息--", msg.SendType)
			//sendP2PMsgDirect(&msg, client)
			SendMsg(&msg, client)
		case <-client.CloseWrite:
			log.Println("gid: ", gid, "uid: ", uid, "发送端关闭")
			return
		}
	}
	fmt.Println("over")
}

func GetClientMsgQueue(gid string, uid string) (chan Msg, error) {
	Client, err := GetClient(uid, gid)
	if err != nil {
		return nil, err
	}
	return Client.MsgQueue.Queue, nil
}

func (client *Client) Close() {
	client.IsClose = true
	client.MsgQueue.IsClose = true

	client.CloseWrite <- 0
	close(client.MsgQueue.Queue)
	client.MsgQueue = nil
	client.Conn.Link.Close()
	client.CloseRequest <- 1
	close(client.CloseRequest)
	close(client.CloseWrite)
	client = nil
}

func (client *Client) DeleteSelf() {
	clients, _ := group.GetClients(client.Gid)
	if clients == nil {
		client.Close()
		return
	}
	clients.DeleteClient(client)
}
