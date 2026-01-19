package server

import (
	"bufio"
	"net"

	"net-cat/domain"
	"net-cat/tools"
)

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

func HandleClient(
	conn net.Conn,
	joinCh chan<- domain.Client, 
	leaveCh chan<- domain.Client, 
	messageCh chan<- domain.Message, 
	UsernameCheckCh chan domain.UsernameCheck,
	limit chan int) {

	// Check if chat room is full
	select {
	case limit <- 1:
	default:
		conn.Write([]byte("CHAT ROOM IS FULL. TRY AGAIN LATER.\n"))
		return
	}
	defer func() { <- limit }()
	defer conn.Close()

	FirstMessage(conn)

	reader := bufio.NewScanner(conn)

	userName := getUsername(conn, reader,UsernameCheckCh)

	joinCh <- domain.Client{Conn: conn, Username: userName}

	for reader.Scan() {
		message := tools.SanitizeInput(reader.Text())
		if message != "" {
			messageCh <- domain.Message{
				Sender:     userName,
				Content:    message,
				SenderConn: conn,
			}
			
		}else{
			Stamp(userName,conn)
		}
	}
	leaveCh <- domain.Client{Conn: conn, Username: userName}
	
}

func getUsername(conn net.Conn, reader *bufio.Scanner,UsernameCheckCh chan domain.UsernameCheck) string {
	conn.Write([]byte("[ENTER YOUR NAME]:"))

	reader.Scan()
	userName := reader.Text()

	userName = CheckUsername(userName, conn, reader,UsernameCheckCh)
	return userName
}

// Check for unique username
func CheckUsername(userName string, conn net.Conn, reader *bufio.Scanner,UsernameCheckCh chan domain.UsernameCheck) string {
	userName = tools.SanitizeInput(userName)
	if userName == "" {
		conn.Write([]byte("INVALID NAME. ENTER ANOTHER NAME:"))
		reader.Scan()
		userName = reader.Text()
	}
	for  {
		replyCh := make(chan bool)
		UsernameCheckCh <- domain.UsernameCheck{
			Username: userName,
			Reply:    replyCh,
		}
		
		if <-replyCh {
			break
		}
		conn.Write([]byte("NAME ALREADY TAKEN. ENTER ANOTHER NAME:"))
		reader.Scan()
		userName = reader.Text()
	}
	return userName
}

func FirstMessage(client net.Conn) {
	client.Write([]byte(WelcomeMessage))
}


