package server

import (
	"bufio"
	"net"
	"strings"
	"time"
)

type Client struct {
	name string
	conn net.Conn
}

// Continuously read messages from client connection:
// - Use bufio.NewScanner(conn)
// - Ignore empty messages
// - Create Message struct
// - Append to server history (protected by mutex)
// - Broadcast message to all clients
// - If connection error:
//   - Remove client from server
//   - Broadcast leave message
//   - Close connection
//   - Exit loop
func (c *Client) ReadLoop(s *Server) {
	defer func() {
		// If loop exits, client is disconnecting

		// Remove client from server
		s.RemoveClient(c.name)

		// Create leave message
		leaveMsg := Message{
			Timestamp: time.Now(),
			Username:  "SYSTEM",
			Content:   c.name + " has left our chat...",
		}

		// Broadcast to remaining clients
		s.Broadcast(leaveMsg)

		// Close connection
		c.conn.Close()
	}()

	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// Ignore empty messages
		if text == "" {
			continue
		}

		msg := Message{
			Timestamp: time.Now(),
			Username:  c.name,
			Content:   text,
		}

		// Broadcast to all clients
		s.Broadcast(msg)
	}

	// On reaching here → scanner stopped (disconnect or error)
}
