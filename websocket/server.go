package websocket

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id string
	color int
	socket *websocket.Conn
}

type Hub struct {
	clients map[*Client]bool
	broadcast chan Message
	register chan *Client
	unregister chan *Client
	mu sync.Mutex
}

type Message struct {
	Sender string `json:"sender"`
	Color string `json:"color"`
	Name string `json:"username"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var hub = Hub {
	clients: make(map[*Client]bool),
	broadcast: make(chan Message),
	register: make(chan *Client),
	unregister: make(chan *Client),
}

func init() {
	go hub.run()
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete (h.clients, client)
				client.socket.Close()
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				if err := client.socket.WriteJSON(message); err != nil {
					client.socket.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

func StartWebSocketServer(port string) {
	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic("Error spinning up WebSocket server: " + err.Error())
	}
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
	client := &Client {
		id: uuid.New().String(),
		color: rand.Intn(360),
		socket: conn,
	}

	hub.register <- client

	defer func() {
		hub.unregister <- client
	}()

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			break
		}
		msg.Sender = client.id
		msg.Color = strconv.Itoa(client.color)
		hub.broadcast <- msg
	}
}



