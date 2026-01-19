package domain

import "net"

type Message struct {
	Sender     string
	SenderConn net.Conn
	Content    string
}
