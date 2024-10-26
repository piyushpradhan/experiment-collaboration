package types

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Color  int
	Socket *websocket.Conn
}

func ValidateClient(c *Client) bool {
	if c.ID == "" || c.Socket == nil {
		return false
	}

	if len(c.ID) < 4 {
		return false
	}

	return true
}
