package domain

import "net"

// Client represents a connected client with its connection and username
type Client struct {
	Conn     net.Conn
	Username string
}
