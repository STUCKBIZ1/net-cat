package utils

import (
	"fmt"
	"time"
)

func formatMessage(name, message string) string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]:%s", currentTime, name, message)
}
