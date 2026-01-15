package utils

import "fmt"

func ChatManager(clientChannel chan Client, messageChannel chan Message) {
	var chatHistory string
	for {
		select {
		case clientInfo := <-clientChannel:
			if clientInfo.conn != nil {
				clients = append(clients, clientInfo)
				jM := clientInfo.name + " has joined a chat \n"
				joinMessage := Message{
					textMessage: jM,
					conn: clientInfo.conn,
				}
				broadCast(joinMessage, clients)
				fmt.Fprint(clientInfo.conn, chatHistory)
			} else {
				removeClient(&clients, clientInfo)
				leftMessage := clientInfo.name + " has left chat\n"
				leftMsgStruct := Message{
					textMessage: leftMessage,
					conn:        nil,
				}
				broadCast(leftMsgStruct, clients)
			}
		case clientMessage := <-messageChannel:
			broadCast(clientMessage, clients)
			chatHistory += clientMessage.textMessage 

		}
	}
}
