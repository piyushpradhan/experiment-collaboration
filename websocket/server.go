package websocket

import (
	"collaboration/types"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var hub = types.Hub{
	Clients:    make(map[*types.Client]bool),
	Broadcast:  make(chan types.Message),
	Register:   make(chan *types.Client),
	Unregister: make(chan *types.Client),
}

func init() {
	go hub.Run()
}

func StartWebSocketServer(port string) {
	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic("Error spinning up WebSocket server: " + err.Error())
	}
	fmt.Println("Starting websocket server on: ", port)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &types.Client{
		ID:     uuid.New().String(),
		Color:  rand.Intn(360),
		Socket: conn,
	}

	hub.Register <- client

	defer func() {
		hub.Unregister <- client
	}()

	for {
		var msg types.Message
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		client.ID = msg.Sender
		msg.Color = strconv.Itoa(client.Color)
		hub.Broadcast <- msg
	}
}
