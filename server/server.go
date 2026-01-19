package server

import (
	"fmt"
	"net"

	"net-cat/tools"
)


func RunServer() (net.Listener, error) {
	// Server logic here
	port, err := tools.CheckPort()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		LogMessage("❌Error starting server:" + fmt.Sprint(err))

		return nil, fmt.Errorf("Error starting server:", err)
	}
	fmt.Println("Listening on the port " + port)
	LogMessage("✅Server started on port " + port)
	return listener, nil
}
