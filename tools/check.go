package tools

import (
	"fmt"
	"os"
	"strings"
)

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

func Atoi(s string) (n int, err error) {
	for _, v := range s {
		if v < '0' || v > '9' {
			return -1, fmt.Errorf("invalid integer")
		}
		n = n*10 + int(v-'0')
	}
	return n, nil
}

func IsPrintable(r rune) bool {
	return (r >= 32 && r <= 126) || (r >= 128 && r <= 255)
}

func SanitizeInput(input string) string {
	sanitized := ""
	for _, r := range input {
		if IsPrintable(r) {
			sanitized += string(r)
		}
	}
	return strings.TrimSpace(sanitized)
}
