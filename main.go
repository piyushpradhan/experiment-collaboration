package main

import (
	"collaboration/api"
	"collaboration/cmd/collaboration"
	"collaboration/storage"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	collaboration.InitializeCollaboration()
	server := api.NewServer(*listenAddr, store)

	fmt.Println("Server running on port: ", *listenAddr)
	log.Fatal(server.Start())
}
