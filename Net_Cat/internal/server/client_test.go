package server

import (
	//"bufio"
	"bytes"
	//"net"
	"sync"
	"testing"
	"time"
)

type TestClient struct {
	*Client
	MockConn *MockConn
}

func NewTestClient(name string) *TestClient {
	mockConn := &MockConn{
		ReadBuffer:  &bytes.Buffer{},
		WriteBuffer: &bytes.Buffer{},
	}
	return &TestClient{
		Client: &Client{
			name: name,
			conn: mockConn,
		},
		MockConn: mockConn,
	}
}

func TestReadLoop_ClientDisconnect(t *testing.T) {
	// Setup server with multiple clients
	client1 := NewTestClient("alice")
	client2 := NewTestClient("bob")
	client3 := NewTestClient("charlie")

	server := &Server{
		Clients: map[string]*Client{
			"alice":   client1.Client,
			"bob":     client2.Client,
			"charlie": client3.Client,
		},
		Messages: []Message{},
		Mu:       sync.Mutex{},
	}

	// Simulate client1 disconnecting (empty read buffer triggers scanner error)
	go client1.ReadLoop(server)

	// Give it time to process disconnect
	time.Sleep(100 * time.Millisecond)

	// Verify client1 was removed
	server.Mu.Lock()
	_, exists := server.Clients["alice"]
	server.Mu.Unlock()

	if exists {
		t.Error("Client was not removed from server after disconnect")
	}

	// Verify leave message was added
	server.Mu.Lock()
	messages := server.Messages
	server.Mu.Unlock()

	found := false
	for _, msg := range messages {
		if msg.Username == "SYSTEM" && msg.Content == "alice has left our chat..." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Leave message not found in message history")
	}

	// Verify other clients still exist
	server.Mu.Lock()
	if len(server.Clients) != 2 {
		t.Errorf("Expected 2 remaining clients, got %d", len(server.Clients))
	}
	if _, exists := server.Clients["bob"]; !exists {
		t.Error("Bob was incorrectly removed")
	}
	if _, exists := server.Clients["charlie"]; !exists {
		t.Error("Charlie was incorrectly removed")
	}
	server.Mu.Unlock()
}
