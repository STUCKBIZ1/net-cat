package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var clients []Client

func HandleConn(conn net.Conn, clientChannel chan Client, messageChannel chan Message) {
	var form string
	var clientName string
	reader := bufio.NewReader(conn)
	logo, _ := os.ReadFile("logolinux.txt")
	if len(clients) > 1 {
		fmt.Fprint(conn, "room chat is full try later")
		conn.Close()
	}
	for {
		if clientName != "" {
			form = formatMessage(clientName, "")
			fmt.Fprint(conn, form)
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
			fmt.Fprint(conn, "Welcome to TCP-Chat!\n")
			fmt.Fprint(conn, string(logo))
			fmt.Fprint(conn, "[ENTER YOUR NAME]:")
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
				}
			} else {
				fmt.Fprint(conn, "cannot use an empty name\n")
			}
		}
	}
}
