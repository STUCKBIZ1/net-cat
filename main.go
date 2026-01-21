package main

import (
	"fmt"

	"net-cat/domain"
	"net-cat/server"
)

func main() {
	// Channels for client management and messaging
	joinCh := make(chan domain.Client)
	leaveCh := make(chan domain.Client)
	messageCh := make(chan domain.Message)
	UsernameCheckCh := make(chan domain.UsernameCheck)
	limit := make(chan int, 10) // Limit to 10 concurrent connections

	// Start the chat manager goroutine
	go server.ChatManager(joinCh, leaveCh, messageCh, UsernameCheckCh)

	// Start the TCP server
	listener, err := server.RunServer()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections
	// Handle each connection in a separate goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connections: ", err)
			server.LogMessage("❌Error accepting connections: " + fmt.Sprint(err))
			return
		}
		go server.HandleClient(conn, joinCh, leaveCh, messageCh, UsernameCheckCh, limit)
	}
}
