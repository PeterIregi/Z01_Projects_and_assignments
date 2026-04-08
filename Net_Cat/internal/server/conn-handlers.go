package server

import (
	"bufio"
	"net"
	"strings"
)

func promptUserForName(name *string, conn net.Conn) {
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	scanner := bufio.NewScanner(conn)

	if !scanner.Scan() { // Input empty
		*name = ""
		return
	}

	*name = scanner.Text()
}

func (s *Server) GetValidUserName(conn net.Conn) string {
	name := ""
	promptUserForName(&name, conn)

	for len(name) == 0 || s.CheckNameExists(&name,conn) {
		promptUserForName(&name, conn)
	}

	return name
}

func (s *Server) CheckNameExists(name *string, conn net.Conn) bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	for _, c := range s.Clients {
		if strings.Trim(c.name," ") == strings.Trim(*name," ") {
			if conn != nil { // Test pass it as nil
				conn.Write([]byte("The name " + *name + " is already taken\n"))
			}
			return true
		}
	}

	return false
}

func (s *Server) SendMsgHistoryToNewConn(conn net.Conn) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	for _, msg := range s.Messages {
		conn.Write([]byte(msg.FormatMessage()))
	}
}
