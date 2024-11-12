package main

import (
	"collaboration/cmd/api"
	"collaboration/cmd/collaboration"
	"flag"
	"fmt"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "the server address")
	flag.Parse()

	go collaboration.InitializeCollaboration()
	api.InitializeApi()

	fmt.Println("Server running on port: ", *listenAddr)
}
