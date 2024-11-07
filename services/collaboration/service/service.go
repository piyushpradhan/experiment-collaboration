package collaboration

import (
	"collaboration/types"
)

type CollaborationService interface {
	RegisterClient(conn *types.Client)
	UnregisterClient(conn *types.Client)
	BroadcastMessage(msg types.Message)
	Run()
}

type collaborationService struct {
	hub *Hub
}

func NewCollaborationService() CollaborationService {
	return &collaborationService{
		hub: NewHub(),
	}
}

func (c *collaborationService) Run() {
	c.hub.Run()
}

func (c *collaborationService) BroadcastMessage(msg types.Message) {
	c.hub.broadcast <- msg
}

func (c *collaborationService) RegisterClient(client *types.Client) {
	c.hub.register <- client
}

func (c *collaborationService) UnregisterClient(client *types.Client) {
	c.hub.unregister <- client
}
