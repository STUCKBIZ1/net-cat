package domain


// UsernameCheck represents a request to check if a username is already taken
// It includes the username to check and a channel to send the result back
type UsernameCheck struct {
	Username string
	Reply    chan bool
}