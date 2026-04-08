package server

import (
	"fmt"
	"net"
	"net-cat/internal/utils"
	"sync"
	"time"
)

type Server struct {
	Clients    map[string]*Client
	Messages   []Message
	Mu         sync.Mutex
	MaxClients int
}

func (s *Server) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return err
	}

	fmt.Println("Listening on the port :", port)
	for {
		conn, _ := listener.Accept()
		fmt.Println(conn.RemoteAddr().String(), " has connected")
		go s.HandleConnection(conn)
	}
}

// Handle new incoming connection:
// - If maxClients reached → reject connection.
// - Send welcome message + ASCII logo.
// - Prompt for name.
// - Read name from connection.
// - Validate non-empty & unique name.
// - Create Client instance.
// - Add client to server (via AddClient).
// - Send previous message history to new client.
// - Broadcast join message to other clients.
// - Start client's ReadLoop in goroutine.
// - Handle client disconnection properly.
func (s *Server) HandleConnection(conn net.Conn) {
	s.Mu.Lock()
	if len(s.Clients) == s.MaxClients {
		msg := fmt.Sprintf("Max number of connections(%v) reached\n", s.MaxClients)
		conn.Write([]byte(msg))
		s.Mu.Unlock()
		conn.Close()
		return
	}
	s.Mu.Unlock()

	// Display linux logo & prompt user for valid user name
	conn.Write([]byte(utils.WelcomeMessage))
	name := s.GetValidUserName(conn)

	// Create client
	newClient := &Client{
		name: name,
		conn: conn,
	}

	// Add client to server
	s.AddClient(newClient)

	// Send previous message history to new client
	s.SendMsgHistoryToNewConn(conn)

	// Broadcast new join message to all
	newJoinMsg := Message{
		Timestamp: time.Now(),
		Username:  "SYSTEM",
		Content:   name + " has joined",
	}
	s.BroadcastExcept(newJoinMsg, name)

	// Start client read loop
	go newClient.ReadLoop(s)
}
