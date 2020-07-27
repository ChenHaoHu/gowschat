package server

import (
	"log"
	"sync"
)

type Clients struct {
	sync.RWMutex
	m   map[string]*Client
	gid string
}

func (clients *Clients) GetClient(uid string) (*Client, error) {
	clients.RLock()
	defer clients.RUnlock()
	client := clients.m[uid]
	return client, nil
}

func (clients *Clients) DeleteClient(client *Client) error {

	if clients == nil {
		return nil
	}
	clients.Lock()
	delete(clients.m, client.Uid)
	clients.Unlock()
	log.Println("client name :", client.Name, " client gid :", client.Gid, " client uid :", client.Uid, " logout ok")
	log.Println("Client : ", client.Gid+" "+client.Uid+" Resource Cleaned")
	client.Close()
	if len(clients.m) == 0 {
		group.DeleteliClents(clients)
	}
	return nil
}

func (clients *Clients) AddClient(client *Client) error {
	clients.Lock()
	clients.m[client.Uid] = client
	clients.Unlock()
	return nil
}

func (clients *Clients) Close() {
	clients = nil
}
