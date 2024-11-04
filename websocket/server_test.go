package websocket

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"collaboration/types"

	"github.com/gorilla/websocket"
)

func TestWebSocketTraffic(t *testing.T) {
	// Start the websocket server in a goroutine
	go StartWebSocketServer(":7071")

	time.Sleep(1 * time.Second)

	url := "ws://localhost:7071/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}

	defer conn.Close()

	// Send a message to the server
	msg := types.Message{
		Sender: "TestUser",
		Name:   "TestUser",
		X:      1.0,
		Y:      2.0,
	}

	err = conn.WriteJSON(msg)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read the broadcasted message back
	var receivedMsg types.Message

	err = conn.ReadJSON(&receivedMsg)
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	// Validate the received message
	if receivedMsg.Sender != msg.Sender {
		t.Errorf("Received message does not match sent message: got %+v, want %+v", receivedMsg, msg)
	}
}

func TestMultipleWebSocketClientsWithCursors(t *testing.T) {
	go StartWebSocketServer(":7071")

	time.Sleep(1 * time.Second)

	url := "ws://localhost:7071/ws"
	numClients := 10
	var wg sync.WaitGroup
	messageCount := make(chan int)

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			conn, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				t.Errorf("Client %d: Failed to connect to websocket server: %v", clientID, err)
				return
			}
			defer conn.Close()

			// Simulate sending cursor position
			for j := 0; j < 5; j++ { // Send 5 cursor updates
				msg := types.Message{
					Sender: fmt.Sprintf("Client%d", clientID),
					X:      float64(clientID + j), // Simulate cursor X position
					Y:      float64(clientID + j + 1), // Simulate cursor Y position
				}
				err = conn.WriteJSON(msg)
				if err != nil {
					t.Errorf("Client %d: Failed to send message: %v", clientID, err)
					return
				}

				// Read the broadcasted message
				var receivedMsg types.Message
				err = conn.ReadJSON(&receivedMsg)
				if err != nil {
					t.Errorf("Client %d: Failed to read message: %v", clientID, err)
					return
				}

				// Signal that a message was received
				messageCount <- 1
				time.Sleep(100 * time.Millisecond) // Simulate delay between cursor updates
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(messageCount)
	}()

	// Count the total messages received
	totalMessages := 0
	for range messageCount {
		totalMessages++
	}

	// Check if the number of messages sent matches the number of messages received
	expectedMessages := numClients * 5 // Each client sends 5 messages
	if totalMessages != expectedMessages {
		t.Errorf("Expected %d messages, but received %d", expectedMessages, totalMessages)
	}
}