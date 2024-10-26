package main

import (
	"collboration/api"
	"collboration/storage"
	"collboration/websocket"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	go websocket.StartWebSocketServer(":7071")
	server := api.NewServer(*listenAddr, store)


	fmt.Println("Server running on port: ", *listenAddr)
	log.Fatal(server.Start())
}
