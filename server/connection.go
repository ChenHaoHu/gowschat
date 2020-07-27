package server

import "github.com/gorilla/websocket"

type Conn struct {
	Link *websocket.Conn
}

func NewConn(wsconn *websocket.Conn) *Conn {
	conn := new(Conn)
	conn.Link = wsconn
	return conn
}

func (conn *Conn) WriteJSON(msg interface{}) error {
	err := conn.Link.WriteJSON(msg)
	return err
}

func (conn *Conn) ReadJSON(v interface{}) error {
	err := conn.Link.ReadJSON(v)
	return err
}

func (conn *Conn) Close() error {
	return conn.Link.Close()
}
