package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)


var frr string



var clients []Client

func HandleConn(conn net.Conn, clientChannel chan Client, messageChannel chan Message) {
	var clientName string
	reader := bufio.NewReader(conn)
	logo, err := os.ReadFile("logolinux.txt")
	if err != nil {

	}
	if len(clients) > 10 {
		fmt.Fprint(conn, "room chat is full try later")
		conn.Close()
	}
	for {
		if clientName != "" {
			frr = formatMessage(clientName, "")
			fmt.Fprint(conn, frr)
			message, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF{
					cl := Client{
						name: clientName,
						conn : nil,
					}
					clientChannel <- cl
					return
				}
			}
			message = formatMessage(clientName, message)
			messageStruct := Message{
				textMessage: message,
				conn:        conn,
			}
			messageChannel <- messageStruct

		} else {
			fmt.Fprint(conn, "welcom to tcp chat\n")
			fmt.Fprint(conn, string(logo))
			fmt.Fprint(conn, "Enter your name : ")
			message, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			message = strings.TrimSpace(message)
			if message != "" {
				clientName = message
				if len(clientName) <= 25 {
					cl := Client{
						name: clientName,
						conn: conn,
					}
					clientChannel <- cl
				} else {
					fmt.Fprint(conn, "cannot use name longer then 25 caracter\n")
					return
				}
			} else {
				fmt.Fprint(conn, "cannot use an empty name\n")
			}
		}
	}
}


