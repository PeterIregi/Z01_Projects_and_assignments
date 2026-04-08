package server

// Add client to server
func (s *Server) AddClient(c *Client) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Clients[c.name] = c
}

// Remove client from server
func (s *Server) RemoveClient(name string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	delete(s.Clients, name)
}

// Broadcast message to all connected clients 
// NOTE: It also adds the message to the server messages slice so DON'T add it outside  the function call
func (s *Server) Broadcast(msg Message) {
	if len(msg.Content) == 0 {
		return
	}

	// Copy clients under lock
	s.Mu.Lock()
	clients := make([]*Client, 0, len(s.Clients))
	for _, c := range s.Clients {
		clients = append(clients, c)
	}

	// Store message
	s.Messages = append(s.Messages, msg)
	s.Mu.Unlock()

	// Send outside lock
	for _, c := range clients {
		if _, err := c.conn.Write([]byte(msg.FormatMessage())); err != nil {
			// Optional: remove broken client
			s.RemoveClient(c.name)
		}
	}
}

// Broadcast message to all connected clients except one
// NOTE: It also adds the message to the server messages slice so DON'T add it outside  the function call
func (s *Server) BroadcastExcept(msg Message, excluded string) {
	if len(msg.Content) == 0 {
		return
	}

	// Copy valid clients under lock
	s.Mu.Lock()
	clients := make([]*Client, 0, len(s.Clients))
	for _, c := range s.Clients {
		if c.name != excluded {
			clients = append(clients, c)
		}
	}

	// Store message
	s.Messages = append(s.Messages, msg)
	s.Mu.Unlock()

	// Send outside lock
	for _, c := range clients {
		if _, err := c.conn.Write([]byte(msg.FormatMessage())); err != nil {
			// Optional: remove broken client
			s.RemoveClient(c.name)
		}
	}
}
