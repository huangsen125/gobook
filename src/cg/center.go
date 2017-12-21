package cg

import (
	"ipc"
	"sync"
	"errors"
	"encoding/json"
)

var _ ipc.Server = &CenterServer{}

type Message struct{
	From string "from"
	To string "to"
	Content string "content"
}

type CenterServer string{
	servers map[string] ipc.Server
	players []*players
	rooms []*Room
	mutex sync.RWMutex
}

func NewCenterServer() *CenterServer  {
	servers := make(map[string] ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{servers:servers,players:players}
}

func (server *CenterServer) addPlayer(params string) error  {
	
	player := NewPlayer()

	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.
}