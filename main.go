package main

import (
	"fmt"
	"net"
	"os"

	"netcat/utils"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) < 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	port := ""
	if len(os.Args) == 1 {
		port = "8989"
	} else {
		port = os.Args[1]
	}

	clientChannel := make(chan utils.Client)
	messageChannel := make(chan utils.Message)

	limit := make(chan int, 10)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	go utils.ChatManager(clientChannel, messageChannel)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go utils.HandleConn(conn, clientChannel, messageChannel, limit)
	}
}
