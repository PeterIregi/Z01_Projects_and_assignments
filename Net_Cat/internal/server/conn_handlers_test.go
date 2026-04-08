package server

import (
	"bytes"
	"sync"
	"testing"
	"time"
)

func TestCheckNameExists(t *testing.T) {
	tests := []struct {
		name        string
		clients     map[string]*Client
		checkName   string
		expected    bool
		description string
	}{
		{
			name:        "Empty server",
			clients:     map[string]*Client{},
			checkName:   "john",
			expected:    false,
			description: "Should return false when server has no clients",
		},
		{
			name: "Name exists in single client server",
			clients: map[string]*Client{
				"john": {
					name: "john",
					conn: &MockConn{},
				},
			},
			checkName:   "john",
			expected:    true,
			description: "Should return true when name exists in single client",
		},
		{
			name: "Name exists in multi-client server",
			clients: map[string]*Client{
				"alice": {
					name: "alice",
					conn: &MockConn{},
				},
				"bob": {
					name: "bob",
					conn: &MockConn{},
				},
				"charlie": {
					name: "charlie",
					conn: &MockConn{},
				},
			},
			checkName:   "bob",
			expected:    true,
			description: "Should return true when name exists among multiple clients",
		},
		{
			name: "Name doesn't exist in multi-client server",
			clients: map[string]*Client{
				"alice": {
					name: "alice",
					conn: &MockConn{},
				},
				"bob": {
					name: "bob",
					conn: &MockConn{},
				},
				"charlie": {
					name: "charlie",
					conn: &MockConn{},
				},
			},
			checkName:   "david",
			expected:    false,
			description: "Should return false when name doesn't exist",
		},
		{
			name: "Case sensitivity - different case",
			clients: map[string]*Client{
				"john": {
					name: "john",
					conn: &MockConn{},
				},
			},
			checkName:   "JOHN",
			expected:    false,
			description: "Should be case sensitive - 'JOHN' ≠ 'john'",
		},
		{
			name: "Name with spaces",
			clients: map[string]*Client{
				"john doe": {
					name: "john doe",
					conn: &MockConn{},
				},
			},
			checkName:   "john doe",
			expected:    true,
			description: "Should handle names with spaces",
		},
		{
			name: "Name with special characters",
			clients: map[string]*Client{
				"user@123": {
					name: "user@123",
					conn: &MockConn{},
				},
			},
			checkName:   "user@123",
			expected:    true,
			description: "Should handle names with special characters",
		},
		{
			name: "Empty string name",
			clients: map[string]*Client{
				"": {
					name: "",
					conn: &MockConn{},
				},
			},
			checkName:   "",
			expected:    true,
			description: "Should handle empty string names",
		},
		{
			name: "Check against nil pointer string",
			clients: map[string]*Client{
				"john": {
					name: "john",
					conn: &MockConn{},
				},
			},
			checkName:   "",
			expected:    false,
			description: "Should handle checking empty string against existing names",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				Clients: tt.clients,
				Mu:      sync.Mutex{},
			}

			result := server.CheckNameExists(&tt.checkName, nil)

			if result != tt.expected {
				t.Errorf("CheckNameExists(%q) = %v, want %v. %s",
					tt.checkName, result, tt.expected, tt.description)
			}
		})
	}
}

func TestServer_SendMsgHistoryToNewConn(t *testing.T) {
	server := &Server{
		Messages: []Message{
			{
				Timestamp: time.Now(),
				Username:  "user1",
				Content:   "Hello",
			},
			{
				Timestamp: time.Now(),
				Username:  "user2",
				Content:   "Hi there",
			},
		},
	}

	mockConn := &MockConn{
		ReadBuffer:  &bytes.Buffer{},
		WriteBuffer: &bytes.Buffer{},
	}

	server.SendMsgHistoryToNewConn(mockConn)

	// Check if messages were written to the connection
	if mockConn.WriteBuffer.Len() == 0 {
		t.Error("Message history was not sent to new connection")
	}
}
