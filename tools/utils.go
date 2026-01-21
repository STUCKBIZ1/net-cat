package tools

import (
	"fmt"
	"os"
	"strings"
)

// CheckPort checks command-line arguments for a port number
// If no port is provided, it defaults to ":8989"
// It returns an error if the provided port is invalid
func CheckPort() (port string, err error) {
	switch len(os.Args) {
	case 1:
		return ":8989", nil
	case 2:
		pnum, err := Atoi(os.Args[1])

		if err == nil && pnum > 0 && pnum < 65536 {
			return ":" + os.Args[1], nil
		} else {
			return "", fmt.Errorf("[USAGE]: ./TCPChat $port")
		}
	default:
		return "", fmt.Errorf("[USAGE]: ./TCPChat $port")
	}
}


// Atoi converts a string to an integer
// It returns an error if the string is not a valid integer
func Atoi(s string) (n int, err error) {
	for _, v := range s {
		if v < '0' || v > '9' {
			return -1, fmt.Errorf("invalid integer")
		}
		n = n*10 + int(v-'0')
	}
	return n, nil
}

// IsPrintable checks if a rune is a printable character
// It returns true for printable characters and false otherwise
// Printable characters are in the ranges 32-126 and 128-255 of the ASCII table
// And the extended ASCII table
// Non-printable characters include control characters and whitespaces other than space
func IsPrintable(r rune) bool {
	return (r >= 32 && r <= 126) || (r >= 128 && r <= 255)
}


// SanitizeInput removes non-printable characters from a string
// It returns the sanitized string with leading and trailing spaces trimmed
func SanitizeInput(input string) string {
	sanitized := ""
	for _, r := range input {
		if IsPrintable(r) {
			sanitized += string(r)
		}
	}
	return strings.TrimSpace(sanitized)
}
