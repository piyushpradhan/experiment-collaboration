package main

import (
	"collaboration/cmd/collaboration"
	"flag"
	"fmt"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "the server address")
	flag.Parse()

	collaboration.InitializeCollaboration()

	fmt.Println("Server running on port: ", *listenAddr)
}
