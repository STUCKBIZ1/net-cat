# Net-Cat (Go)

A simple **TCP chat server** written in **Go**.
Multiple clients can connect, choose a unique username, send messages, and receive chat history.

---

## Features

* ✅ TCP server using Go `net` package
* 👥 Multiple clients (concurrent connections)
* 🔒 Limit connections (max 10 clients)
* 🧵 Goroutines + Channels for concurrency
* 🆔 Unique username check
* 📜 Chat history for new clients
* 📝 Server logs (`chat.log`)
* 🧹 Input sanitization
* 🎨 Colored system messages (join / leave)

---

## Project Structure

```yaml
.
├── main.go
├── domain/
│   ├── client.go
│   ├── message.go
│   └── username_check.go
├── server/
│   ├── chat_manager.go
│   ├── handler.go
│   └── server.go
├── tools/
│   └── utils.go
├── chat.log
└── README.md
```

---

## How It Works (High Level)

* **main.go**

  * Creates channels (join, leave, messages, username check)
  * Starts `ChatManager` in a goroutine
  * Accepts TCP connections
  * Each client is handled in its own goroutine

* **ChatManager** (the brain 🧠)

  * Manages connected clients
  * Broadcasts messages
  * Handles join / leave events
  * Checks username availability

* **HandleClient**

  * Limits connections
  * Asks for username
  * Reads messages from the client
  * Sends messages through channels

---

## Channels Used

| Channel           | Purpose                  |
| ----------------- | ------------------------ |
| `joinCh`          | New client joined        |
| `leaveCh`         | Client left              |
| `messageCh`       | Chat messages            |
| `UsernameCheckCh` | Check unique usernames   |
| `limit`           | Limit concurrent clients |

---

## Installation

### Requirements

* Go 1.20+

### Clone

```bash
git clone https://learn.zone01oujda.ma/git/maadlani/net-cat.git
cd net-cat
```

---

## Run the Server

### Default port (8989)

```bash
go run .
```

### Custom port

```css
go run . 2525
```

---

## Connect as a Client

Using **netcat**:

```css
nc localhost 8989
```

You will see:

```
[ENTER YOUR NAME]:
```

---

## Chat Rules

* Username must be unique
* Empty messages are ignored
* Non-printable characters are removed
* Max 10 clients connected at the same time

---

## Logging

* All server events and messages are saved in:

```yaml
chat.log
```

---

## Error Handling

* Port already in use
* Invalid port number
* Chat room full
* Invalid username

---

## Example Output

```css
[2026-01-19 15:19:44][System]: M has joined our chat...
[2026-01-19 15:20:09][A]: hey
[2026-01-19 15:20:18][M]: test1
```

---

## Technologies

* Go
* TCP Networking
* Goroutines
* Channels

---

## Authors

**Abdelali**
Zone01 / 1337 Student

**Meryam**
Zone01 Student

---

## License

This project is for **educational purposes**.
