package domain

import "net"

// Message represents a chat message sent by a client
// It includes the sender's username, connection, and the message content
type Message struct {
	Sender     string
	SenderConn net.Conn
	Content    string
}
