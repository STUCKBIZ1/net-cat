package server

import (
	"bufio"
	"fmt"
	"net"

	"net-cat/domain"
	"net-cat/tools"
)

// Welcome banner sent to the client on connection
const WelcomeMessage = `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    ` + "`" + `.       | ` + "`" + `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     ` + "`" + `-'       ` + "`" + `--'
`
const MaxMessageLength = 3*1024

// HandleClient manages the interaction with a connected client
// It handles username assignment, message reading, and disconnection
// It also enforces a limit on concurrent connections using the limit channel
// On disconnection, it notifies the chat manager via the leaveCh channel
// It sanitizes input messages and checks for maximum message length
func HandleClient(
	conn net.Conn,
	joinCh chan<- domain.Client,
	leaveCh chan<- domain.Client,
	messageCh chan<- domain.Message,
	UsernameCheckCh chan domain.UsernameCheck,
	limit chan int,
) {
	// Check if chat room is full
	// If full, notify client and close connection
	defer conn.Close()
	select {
	case limit <- 1:
	default:
		conn.Write([]byte("CHAT ROOM IS FULL. TRY AGAIN LATER.\n"))
		return
	}
	defer func() { <-limit }()

	FirstMessage(conn)

	reader := bufio.NewScanner(conn)

	userName := getUsername(conn, reader, UsernameCheckCh)

	joinCh <- domain.Client{Conn: conn, Username: userName}
	// Read messages from the client
	for reader.Scan() {
		message := tools.SanitizeInput(reader.Text())
		if len(message)> MaxMessageLength {
			conn.Write([]byte("MESSAGE TOO LONG. MAX LENGTH IS " + fmt.Sprint(MaxMessageLength) + " CHARACTERS.\n"))
			Stamp(userName,conn)
			continue
		}
		if message != ""  {
			messageCh <- domain.Message{
				Sender:     userName,
				Content:    message,
				SenderConn: conn,
			}
		} else {
			Stamp(userName, conn)
		}
	}
	leaveCh <- domain.Client{Conn: conn, Username: userName}
}

// Prompt client for username and validate it
func getUsername(conn net.Conn, reader *bufio.Scanner, UsernameCheckCh chan domain.UsernameCheck) string {
	conn.Write([]byte("[ENTER YOUR NAME]:"))

	reader.Scan()
	userName := reader.Text()

	userName = CheckUsername(userName, conn, reader, UsernameCheckCh)
	return userName
}

// Check for unique username
// Recursively prompt for a new username if invalid or taken
func CheckUsername(userName string, conn net.Conn, reader *bufio.Scanner, UsernameCheckCh chan domain.UsernameCheck) string {
	userName = tools.SanitizeInput(userName)
	if userName == "" || len(userName) > 25 {
		conn.Write([]byte("INVALID NAME. ENTER ANOTHER NAME:"))
		reader.Scan()
		userName = reader.Text()
		userName = CheckUsername(userName, conn, reader, UsernameCheckCh)
	}
	for {
		replyCh := make(chan bool)
		UsernameCheckCh <- domain.UsernameCheck{
			Username: userName,
			Reply:    replyCh,
		}

		if <-replyCh {
			break
		}
		conn.Write([]byte("NAME ALREADY TAKEN. ENTER ANOTHER NAME:"))
		if reader.Scan() {

			userName = reader.Text()
			userName = CheckUsername(userName, conn, reader, UsernameCheckCh)
		}
	}
	return userName
}

// Send the welcome message to the client
func FirstMessage(client net.Conn) {
	client.Write([]byte(WelcomeMessage))
}
