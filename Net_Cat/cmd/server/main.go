package main

import (
	"fmt"
	s "net-cat/internal/server"
	"net-cat/internal/utils"
	"os"
)

func main() {
	//  CLI parsing logic
	port, err := utils.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Instantiate Server
	server := &s.Server{
		Clients:    make(map[string]*s.Client),
		Messages:   []s.Message{},
		MaxClients: 10,
	}

	// Start server
	server.Start(port)
}
