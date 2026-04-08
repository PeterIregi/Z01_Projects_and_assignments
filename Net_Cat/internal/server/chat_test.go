package server

import (
	"bytes"
	"net"
	"testing"
	"time"
	//"sync"
)

/*
	type Server struct {
		Clients    map[string]*Client
		Messages   []Message
		Mu         sync.Mutex
		MaxClients int
	}

	type Client struct {
		name string
		conn net.Conn
	}
*/
type MockConn struct {
	ReadBuffer  *bytes.Buffer
	WriteBuffer *bytes.Buffer
	CloseCalled bool
}

//type Message struct {
//	Timestamp time.Time
//	Username  string
//	Content   string
//}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.ReadBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.WriteBuffer.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return &net.TCPAddr{}
}

func (m *MockConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{}
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestAddClient(t *testing.T) {
	server := &Server{
		Clients: make(map[string]*Client),
	}

	client := &Client{
		name: "testuser",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}

	server.AddClient(client)

	if len(server.Clients) != 1 {
		t.Errorf("Expected 1 client, got %d", len(server.Clients))
	}

	if _, exists := server.Clients["testuser"]; !exists {
		t.Error("Client was not added to the map with correct key")
	}
}

func TestRemoveClient(t *testing.T) {
	server := &Server{
		Clients: make(map[string]*Client),
	}

	client := &Client{
		name: "testuser",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}

	server.AddClient(client)
	server.RemoveClient("testuser")

	if len(server.Clients) != 0 {
		t.Errorf("Expected 0 clients, got %d", len(server.Clients))
	}
}

func TestBroadcast(t *testing.T) {
	server := &Server{
		Clients: make(map[string]*Client),
	}

	// Create mock connections for clients
	client1 := &Client{
		name: "user1",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}
	client2 := &Client{
		name: "user2",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}

	server.AddClient(client1)
	server.AddClient(client2)

	msg := Message{
		Timestamp: time.Now(),
		Username:  "user1",
		Content:   "Hello everyone!",
	}

	server.Broadcast(msg)

	// Check if both clients received the message
	mockConn1 := client1.conn.(*MockConn)
	mockConn2 := client2.conn.(*MockConn)

	if mockConn1.WriteBuffer.Len() == 0 {
		t.Error("Client 1 did not receive broadcast message")
	}
	if mockConn2.WriteBuffer.Len() == 0 {
		t.Error("Client 2 did not receive broadcast message")
	}
}

func TestServer_BroadcastExcept(t *testing.T) {
	server := &Server{
		Clients: make(map[string]*Client),
	}

	// Create mock connections for clients
	client1 := &Client{
		name: "user1",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}
	client2 := &Client{
		name: "user2",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}
	client3 := &Client{
		name: "user3",
		conn: &MockConn{WriteBuffer: &bytes.Buffer{}},
	}

	server.AddClient(client1)
	server.AddClient(client2)
	server.AddClient(client3)

	msg := Message{
		Timestamp: time.Now(),
		Username:  "SYSTEM",
		Content:   "user2 has joined",
	}

	server.BroadcastExcept(msg, "user2")

	// Check if only user2 didn't receive the message
	mockConn1 := client1.conn.(*MockConn)
	mockConn2 := client2.conn.(*MockConn)
	mockConn3 := client3.conn.(*MockConn)

	if mockConn1.WriteBuffer.Len() == 0 {
		t.Error("Client 1 did not receive broadcast message")
	}
	if mockConn2.WriteBuffer.Len() != 0 {
		t.Error("Client 2 received message but should have been excluded")
	}
	if mockConn3.WriteBuffer.Len() == 0 {
		t.Error("Client 3 did not receive broadcast message")
	}
}
