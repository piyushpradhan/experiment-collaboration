package collaboration

import (
	"collaboration/types"
	"sync"
)

type Hub struct {
	clients map[*types.Client]bool
	broadcast chan types.Message
	register chan *types.Client
	unregister chan *types.Client
	mu sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*types.Client]bool),
		broadcast: make(chan types.Message),
		register: make(chan *types.Client),
		unregister: make(chan *types.Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: 
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Socket.Close()
			}
			h.mu.Unlock()
		case msg := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				client.Socket.WriteJSON(msg)
			}
		}
	}
}
