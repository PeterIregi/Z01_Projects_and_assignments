package utils

import (
	"errors"
	"strings"
)

const defaultPort = "8989"

// ParseArgs parses CLI arguments and returns the port.
// Rules:
// - No args        -> return default port (8989)
// - One arg        -> validate numeric port
// - More than one  -> return usage error
func ParseArgs(args []string) (string, error) {
	// No port provided → use default
	if len(args) == 0 {
		return defaultPort, nil
	}

	// More than one argument → invalid usage
	if len(args) > 1 {
		return "", errors.New("[USAGE]: ./TCPChat $port")
	}

	port := strings.TrimSpace(args[0])

	// Empty string not allowed
	if port == "" {
		return "", errors.New("[USAGE]: ./TCPChat $port")
	}

	// Validate numeric port (only digits allowed)
	for _, r := range port {
		if r < '0' || r > '9' {
			return "", errors.New("[USAGE]: ./TCPChat $port")
		}
	}

	return port, nil
}
