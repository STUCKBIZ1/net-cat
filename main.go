package main

import (
	"fmt"

	"net-cat/domain"
	"net-cat/server"
)

func main() {
	joinCh := make(chan domain.Client)
	leaveCh := make(chan domain.Client)
	messageCh := make(chan domain.Message)
	UsernameCheckCh := make(chan domain.UsernameCheck)
	limit := make(chan int, 10) // Limit to 10 concurrent connections

	go server.ChatManager(joinCh, leaveCh, messageCh, UsernameCheckCh)

	listener, err := server.RunServer()
	if err != nil {
		fmt.Println(err)
		return
	}

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
