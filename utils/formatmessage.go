package utils

import (
	"fmt"
	"strings"
	"time"
)

func formatMessage(name, message string) string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if strings.TrimSpace(message) == "" {
		return fmt.Sprintf("[%s][%s]: ", currentTime, name)
	}
	return fmt.Sprintf("[%s][%s]: %s", currentTime, name, message)
}
