package server

import (
	"testing"
	"time"
)

func TestFormatMessage(t *testing.T) {
	testTime := time.Date(2026, 3, 10, 14, 30, 45, 0, time.UTC)
	msg := Message{
		Timestamp: testTime,
		Username:  "daniel",
		Content:   "hello world",
	}

	expected := "[2026-03-10 14:30:45][daniel]:hello world\n"
	result := msg.FormatMessage()

	if result != expected {
		t.Errorf("expected %s but got %s", expected, result)
	}
}
