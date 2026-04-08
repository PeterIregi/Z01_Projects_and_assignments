package utils

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedPort  string
		expectedError error
	}{
		{
			name:          "No arguments - returns default port",
			args:          []string{},
			expectedPort:  "8989",
			expectedError: nil,
		},
		{
			name:          "Valid port - single digit",
			args:          []string{"7"},
			expectedPort:  "7",
			expectedError: nil,
		},
		{
			name:          "Valid port - multiple digits",
			args:          []string{"8080"},
			expectedPort:  "8080",
			expectedError: nil,
		},
		{
			name:          "Valid port with whitespace",
			args:          []string{"  9090  "},
			expectedPort:  "9090",
			expectedError: nil,
		},
		{
			name:          "More than one argument - error",
			args:          []string{"8080", "9090"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Empty string argument",
			args:          []string{""},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Whitespace only argument",
			args:          []string{"   "},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Non-numeric port - letters",
			args:          []string{"abc"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Non-numeric port - alphanumeric",
			args:          []string{"8080a"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Non-numeric port - special characters",
			args:          []string{"80-80"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Non-numeric port - decimal point",
			args:          []string{"8080.0"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Port with leading zeros",
			args:          []string{"00080"},
			expectedPort:  "00080",
			expectedError: nil,
		},
		{
			name:          "Very long port number",
			args:          []string{"1234567890"},
			expectedPort:  "1234567890",
			expectedError: nil,
		},
		{
			name:          "Port with negative sign",
			args:          []string{"-8080"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Port with plus sign",
			args:          []string{"+8080"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "Multiple arguments with valid ports",
			args:          []string{"8080", "9090", "7070"},
			expectedPort:  "",
			expectedError: errors.New("[USAGE]: ./TCPChat $port"),
		},
		{
			name:          "One argument with multiple spaces",
			args:          []string{"   8080   "},
			expectedPort:  "8080",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port, err := ParseArgs(tt.args)

			// Check error
			if tt.expectedError == nil && err != nil {
				t.Errorf("ParseArgs() returned unexpected error: %v", err)
			}
			if tt.expectedError != nil && err == nil {
				t.Errorf("ParseArgs() expected error %v, got nil", tt.expectedError)
			}
			if tt.expectedError != nil && err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("ParseArgs() error = %v, want %v", err, tt.expectedError)
			}

			// Check port
			if port != tt.expectedPort {
				t.Errorf("ParseArgs() port = %v, want %v", port, tt.expectedPort)
			}
		})
	}
}

func TestParseArgsEdgeCases(t *testing.T) {
	// Test with nil slice (different from empty slice)
	t.Run("Nil slice", func(t *testing.T) {
		var nilSlice []string
		port, err := ParseArgs(nilSlice)

		if err != nil {
			t.Errorf("ParseArgs() with nil slice returned error: %v", err)
		}
		if port != "8989" {
			t.Errorf("ParseArgs() with nil slice returned %v, want 8989", port)
		}
	})

	// Test with very large port number (string length)
	t.Run("Extremely large port number", func(t *testing.T) {
		largePort := make([]byte, 10000)
		for i := range largePort {
			largePort[i] = '9'
		}
		args := []string{string(largePort)}

		port, err := ParseArgs(args)

		if err != nil {
			t.Errorf("ParseArgs() with large port returned error: %v", err)
		}
		if len(port) != 10000 {
			t.Errorf("ParseArgs() with large port returned length %d, want 10000", len(port))
		}
	})

	// Test with Unicode digits
	t.Run("Unicode digits", func(t *testing.T) {
		// These are Unicode digits but not ASCII digits
		unicodeDigits := []string{"①", "٢", "۳"}
		for _, digit := range unicodeDigits {
			args := []string{digit}
			port, err := ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() with Unicode digit %q should return error, got port %v", digit, port)
			}
			if port != "" {
				t.Errorf("ParseArgs() with Unicode digit %q returned port %v, want empty string", digit, port)
			}
		}
	})

	// Test with mixed whitespace and valid port
	t.Run("Mixed whitespace patterns", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected string
		}{
			{"\t8080", "8080"},
			{"8080\n", "8080"},
			{"\n8080\t", "8080"},
			{" \t 8080 \n\t ", "8080"},
		}

		for _, test := range inputs {
			args := []string{test.input}
			port, err := ParseArgs(args)

			if err != nil {
				t.Errorf("ParseArgs() with input %q returned error: %v", test.input, err)
			}
			if port != test.expected {
				t.Errorf("ParseArgs() with input %q = %v, want %v", test.input, port, test.expected)
			}
		}
	})
}

func TestParseArgsPortValidation(t *testing.T) {
	// Test various port ranges (though ParseArgs doesn't validate range, just digits)
	validPorts := []string{
		"0",
		"1",
		"80",
		"443",
		"1024",
		"49151",
		"65535",
		"65536", // Above valid port range but still numeric - should pass digit check
		"99999",
	}

	for _, port := range validPorts {
		t.Run("Valid port digits: "+port, func(t *testing.T) {
			args := []string{port}
			result, err := ParseArgs(args)

			if err != nil {
				t.Errorf("ParseArgs() with port %q returned error: %v", port, err)
			}
			if result != port {
				t.Errorf("ParseArgs() with port %q = %v, want %v", port, result, port)
			}
		})
	}

	// Test invalid port formats (should fail digit check)
	invalidPorts := []struct {
		port   string
		reason string
	}{
		{"80.80", "contains decimal point"},
		{"80,80", "contains comma"},
		{"80-80", "contains hyphen"},
		{"80 80", "contains space"},
		{"0x50", "hexadecimal"},
		{"0b1010", "binary"},
		{"80a", "alphanumeric"},
		{"a80", "alphanumeric prefix"},
		{"8O80", "letter O instead of zero"},
		{"", "empty string"},
		{"   ", "only whitespace"},
	}

	for _, test := range invalidPorts {
		t.Run("Invalid port: "+test.port, func(t *testing.T) {
			args := []string{test.port}
			port, err := ParseArgs(args)

			if err == nil {
				t.Errorf("ParseArgs() with %q (%s) should return error, got port %v",
					test.port, test.reason, port)
			}
			if err != nil && err.Error() != "[USAGE]: ./TCPChat $port" {
				t.Errorf("ParseArgs() with %q returned wrong error: %v", test.port, err)
			}
			if port != "" {
				t.Errorf("ParseArgs() with %q returned port %v, want empty string", test.port, port)
			}
		})
	}
}

func TestParseArgsErrorMessages(t *testing.T) {
	// Verify error messages are exactly as specified
	tests := []struct {
		name          string
		args          []string
		expectedError string
	}{
		{
			name:          "Multiple args error",
			args:          []string{"8080", "9090"},
			expectedError: "[USAGE]: ./TCPChat $port",
		},
		{
			name:          "Empty string error",
			args:          []string{""},
			expectedError: "[USAGE]: ./TCPChat $port",
		},
		{
			name:          "Non-numeric error",
			args:          []string{"abc"},
			expectedError: "[USAGE]: ./TCPChat $port",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseArgs(tt.args)

			if err == nil {
				t.Errorf("ParseArgs() expected error, got nil")
			} else if err.Error() != tt.expectedError {
				t.Errorf("ParseArgs() error = %q, want %q", err.Error(), tt.expectedError)
			}
		})
	}
}

// Benchmark tests
func BenchmarkParseArgs(b *testing.B) {
	// Benchmark with no args
	b.Run("No args", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseArgs([]string{})
		}
	})

	// Benchmark with valid port
	b.Run("Valid port", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseArgs([]string{"8080"})
		}
	})

	// Benchmark with invalid port
	b.Run("Invalid port", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseArgs([]string{"abc"})
		}
	})

	// Benchmark with multiple args
	b.Run("Multiple args", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseArgs([]string{"8080", "9090"})
		}
	})

	// Benchmark with port containing whitespace
	b.Run("Port with whitespace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ParseArgs([]string{"  8080  "})
		}
	})
}

// Example tests for documentation
func ExampleParseArgs() {
	// No arguments
	port, _ := ParseArgs([]string{})
	fmt.Println(port)

	// Single valid argument
	port, _ = ParseArgs([]string{"8080"})
	fmt.Println(port)

	// Invalid argument
	_, err := ParseArgs([]string{"abc"})
	fmt.Println(err)

	// Output:
	// 8989
	// 8080
	// [USAGE]: ./TCPChat $port
}

// Test with different package import scenarios
func TestParseArgsPackageUsage(t *testing.T) {
	// Verify the function works with typical command line args
	// Simulate os.Args[1:] behavior
	t.Run("Simulate os.Args[1:]", func(t *testing.T) {
		// Simulate: ./TCPChat
		args := []string{}
		port, err := ParseArgs(args)
		if err != nil || port != "8989" {
			t.Errorf("Failed with no args: got %v, %v", port, err)
		}

		// Simulate: ./TCPChat 8080
		args = []string{"8080"}
		port, err = ParseArgs(args)
		if err != nil || port != "8080" {
			t.Errorf("Failed with single arg: got %v, %v", port, err)
		}

		// Simulate: ./TCPChat 8080 extra
		args = []string{"8080", "extra"}
		_, err = ParseArgs(args)
		if err == nil {
			t.Error("Expected error with multiple args, got nil")
		}
	})
}
