package utils

import "fmt"

func broadCast(clientMessage Message, clients []Client) {
	var msg string
	for _, client := range clients {
		if client.conn != clientMessage.conn {
			msg = formatMessage(client.name, "")
			fmt.Fprint(client.conn, "\n"+clientMessage.textMessage)
			fmt.Fprint(client.conn, msg)
		}
	}
}
