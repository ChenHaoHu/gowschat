package server

import (
	"errors"
	"log"
	"sync"
)

type Group struct {
	sync.RWMutex
	g map[string]*Clients
}

func GetAllOnLineClient() map[string]*Clients {
	return group.g
}

func GetGroupOnLineClient(gid string) map[string]*Client {

	if group.g[gid] == nil {
		return nil
	}

	return group.g[gid].m
}

func (group *Group) AddClients(clients *Clients) error {
	group.Lock()
	defer group.Unlock()
	group.g[clients.gid] = clients
	return nil
}

func (group *Group) GetClients(groupId string) (*Clients, error) {
	group.RLock()
	defer group.RUnlock()
	clients := group.g[groupId]
	return clients, nil
}

func (group *Group) DeleteliClents(clients *Clients) error {
	group.Lock()
	defer group.Unlock()
	delete(group.g, clients.gid)
	clients.Close()
	return nil
}

func (group *Group) GetClientsNotNil(groupId string) (*Clients, error) {
	group.RLock()
	clients := group.g[groupId]
	group.RUnlock()
	if clients == nil {
		clients = &Clients{
			m:   make(map[string]*Client),
			gid: groupId,
		}
		group.Lock()
		group.g[groupId] = clients
		group.Unlock()
	}
	return clients, nil
}

func (group *Group) RegisterClientToGroup(client *Client) {
	clients, _ := group.GetClientsNotNil(client.Gid)
	oldClient, _ := clients.GetClient(client.Uid)
	if oldClient != nil {
		//close conn
		// sendN2PMsg(&Msg{Uid: oldClient.Uid, ToUid: oldClient.Uid, Gid: gid,
		// 	Msg: "Your account is repeatedly logged in in this group", SendType: N2P})
		clients.DeleteClient(oldClient)
		// oldClient.Conn.Close()
	}
	clients.AddClient(client)
	group.AddClients(clients)

	log.Println("client name :", client.Name, " client gid :", client.Gid, " client uid :", client.Uid, " login ok")
}

func GetClient(uid string, gid string) (*Client, error) {

	clients, _ := group.GetClients(gid)

	if clients == nil {
		return nil, errors.New("no group")
	}

	client, _ := clients.GetClient(uid)

	if client != nil {
		return client, nil
	} else {
		return nil, errors.New(gid + "---" + uid + "----no Client")
	}

}
