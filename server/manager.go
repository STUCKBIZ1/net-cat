package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"net-cat/domain"
)

var historyMessages []string

// ChatManager handles client join/leave events and message broadcasting
// It maintains a list of connected clients and processes incoming messages
// It also checks for username availability
// and broadcasts system messages for joins/leaves
// It saves chat history and sends it to new clients upon joining
func ChatManager(
	joinCh chan domain.Client, 
	leaveCh chan domain.Client, 
	messageCh chan domain.Message, 
	UsernameCheckCh chan domain.UsernameCheck) {

	Clients := make(map[net.Conn]domain.Client)

	for {
		select {
		case client := <-joinCh:
			SendHistory(client.Conn)
			Clients[client.Conn] = client
			Stamp(client.Username, client.Conn)
			joined(client.Username, client.Conn, Clients)
		case client := <-leaveCh:
			left(client.Username, client.Conn, Clients)
			delete(Clients, client.Conn)
		case message := <-messageCh:
			broadcast(stampMessage(message.Sender)+message.Content, message.SenderConn, Clients)
			Stamp(message.Sender, message.SenderConn)
			
		case req := <-UsernameCheckCh:
			available := true
			for _, client := range Clients {
				if client.Username == req.Username {
					available = false
					break
				}
			}
			req.Reply <- available
		}
	}
}

// Function to broadcast messages to all connected clients except the sender
// It also logs the message and saves it to chat history
func broadcast(msg string, sender net.Conn, Clients map[net.Conn]domain.Client) {
	for client := range Clients {
		if client != sender {
			client.Write([]byte("\r\033[2K" + msg + "\n"))
			client.Write([]byte(stampMessage(Clients[client].Username)))
		}
	}
	LogMessage("✅" + msg)
	SaveHistory(msg)
}

// LogMessage appends a message to the chat log file 
// It creates the file if it doesn't exist
func LogMessage(message string) {
	file, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	file.WriteString(message + "\n")
}

// SaveHistory appends a message to the in-memory chat history slice
// It returns the updated history slice
func SaveHistory(msg string) []string {
	historyMessages = append(historyMessages, msg)
	return historyMessages
}

// SendHistory sends the saved chat history to a newly connected client
func SendHistory(conn net.Conn) {
	if historyMessages == nil {
		return
	}
	conn.Write([]byte("---- Chat History ----\n"))
	for _, line := range historyMessages {
		conn.Write([]byte(line + "\n"))
	}
	conn.Write([]byte("---- End of History ----\n"))
}

// Broadcast system message when a user joins
func joined(userName string, conn net.Conn, Clients map[net.Conn]domain.Client) {
	broadcast("\033[32m"+stampMessage("System")+"\033[0m"+userName+"\033[32m has joined our chat...\033[0m", conn, Clients)
}

// Broadcast system message when a user leaves
func left(userName string, conn net.Conn, Clients map[net.Conn]domain.Client) {
	broadcast("\033[31m"+stampMessage("System")+"\033[0m"+userName+"\033[31m has left our chat...\033[0m", conn, Clients)
}

// Send timestamp prompt to the client 
func Stamp(userName string, conn net.Conn) {
	conn.Write([]byte(stampMessage(userName)))
}

// Create a timestamped message prefix
func stampMessage(userName string) string {
	return "\033[0m[\033[94m" + time.Now().Format(time.DateTime) + "\033[0m][\033[93m" + userName + "\033[0m]:"
}
