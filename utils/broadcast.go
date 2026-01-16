package utils

import "fmt"

func broadCast(clientMessage Message, clients []Client) {
	for _, client := range clients {
		if client.conn != nil && client.conn != clientMessage.conn {
			fmt.Fprint(client.conn, clientMessage.textMessage)
		}
	}
}
