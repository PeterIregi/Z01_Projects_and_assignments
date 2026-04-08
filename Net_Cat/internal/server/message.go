package server

import (
	"fmt"
	"time"
)

type Message struct {
	Timestamp time.Time
	Username  string
	Content   string
}

// Format message as:
// [YYYY-MM-DD HH:MM:SS][username]:content
// Use m.Timestamp.Format("2006-01-02 15:04:05")
func (m *Message) FormatMessage() string {
	return fmt.Sprintf(
		"[%s][%s]:%s",
		m.Timestamp.Format("2006-01-02 15:04:05"),
		m.Username,
		m.Content,
	) + "\n"
}
