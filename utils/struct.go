package utils

import "net"

type Client struct {
	name string
	conn net.Conn
}

type Message struct {
	textMessage string
	conn        net.Conn
}
