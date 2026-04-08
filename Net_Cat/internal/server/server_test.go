package server

import (
	"bytes"
	"net"
	"testing"
)

func TestServer_HandleConnection(t *testing.T) {
	tests := []struct {
		name          string
		server        *Server
		setupMock     func() net.Conn
		expectClosed  bool
		expectMessage string
	}{
		{
			name: "Reject connection when max clients reached",
			server: &Server{
				Clients:    make(map[string]*Client),
				MaxClients: 1,
			},
			setupMock: func() net.Conn {
				// Fill the server with one client
				mockConn := &MockConn{
					ReadBuffer:  bytes.NewBuffer([]byte("testuser\n")),
					WriteBuffer: &bytes.Buffer{},
				}
				return mockConn
			},
			expectClosed:  true,
			expectMessage: "Max number of connections(1) reached\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add a client to reach max capacity
			if tt.name == "Reject connection when max clients reached" {
				existingClient := &Client{
					name: "existing",
					conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
				}
				tt.server.Clients["existing"] = existingClient
			}

			conn := tt.setupMock()
			mockConn := conn.(*MockConn)

			tt.server.HandleConnection(conn)

			// Check if connection was closed (in this mock, we just check if write buffer contains the rejection message)
			if tt.expectClosed {
				if mockConn.WriteBuffer.String() != tt.expectMessage {
					t.Errorf("Expected rejection message %q, got %q", tt.expectMessage, mockConn.WriteBuffer.String())
				}
			}
		})
	}
}

func TestServer_ConcurrentAccess(t *testing.T) {
	server := &Server{
		Clients: make(map[string]*Client),
	}

	// Test concurrent client additions
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(id int) {
			client := &Client{
				name: string(rune('A' + id)),
				conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
			}
			server.AddClient(client)
			done <- true
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}

	server.Mu.Lock()
	if len(server.Clients) != 10 {
		t.Errorf("Expected 10 clients after concurrent additions, got %d", len(server.Clients))
	}
	server.Mu.Unlock()
}
