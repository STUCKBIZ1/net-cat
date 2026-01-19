package domain

type UsernameCheck struct {
	Username string
	Reply    chan bool
}