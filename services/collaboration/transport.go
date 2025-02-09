package collaboration

import (
	collaboration "collaboration/services/collaboration/service"
	"collaboration/types"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func StartWebSocketServer(svc collaboration.CollaborationService, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(svc, w, r)
	})
	fmt.Println("Starting websocket server on: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic("Error starting WebSocket server: " + err.Error())
	}
}

func handleConnections(svc collaboration.CollaborationService, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO: Add log
		return
	}

	client := &types.Client{
		ID:     uuid.New().String(),
		Color:  rand.Intn(360),
		Socket: conn,
	}

	svc.RegisterClient(client)

	defer func() {
		svc.UnregisterClient(client)
	}()

	for {
		var msg types.Message
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		client.ID = msg.Sender
		msg.Color = strconv.Itoa(client.Color)
		svc.BroadcastMessage(msg)
	}
}
