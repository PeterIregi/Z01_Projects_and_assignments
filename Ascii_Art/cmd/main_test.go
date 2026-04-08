package main

import (
	"bytes"
	"os"
	ap "ascii_art/pkg"
	"testing"
	"fmt"
	"strings"
)


// capture stdout helper
func captureOutput(t *testing.T, f func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe error: %v", err)
	}

	os.Stdout = w
	f()
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("read error: %v", err)
	}

	return buf.String()
}

func TestPrintAscii(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "newline only",
			input:    "\n",
			expected: "\n",
		},
		{
			name:  "Hello",
			input: "Hello",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
		{
			name:  "hello lowercase",
			input: "hello",
			expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
		},
		{
			name:  "HeLlO mixed",
			input: "HeLlO",
			expected: ` _    _          _        _    ____   
| |  | |        | |      | |  / __ \  
| |__| |   ___  | |      | | | |  | | 
|  __  |  / _ \ | |      | | | |  | | 
| |  | | |  __/ | |____  | | | |__| | 
|_|  |_|  \___| |______| |_|  \____/  
                                      
                                      
`,
		},
		{
			name:  "Hello There",
			input: "Hello There",
			expected: ` _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 
                                                                             
                                                                             
`,
		},
		{
			name:  "numbers in input",
			input: "1Hello 2There",
			expected: `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`,
		},
		{
			name:  "symbols",
			input: "{Hello There}",
			expected: `   __  _    _          _   _                 _______   _                           __    
  / / | |  | |        | | | |               |__   __| | |                          \ \   
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ 
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  
  \_\                                                                              /_/   
                                                                                         
`,
		},
		{
			name:  "newline separation",
			input: "Hello\nThere",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
		},
		{
			name:  "double newline",
			input: "Hello\n\nThere",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
		},
	}
	banner := ap.LoadBanner("../standard.txt")

	for _, tt := range tests {


		t.Run(tt.name, func(t *testing.T) {
			out := captureOutput(t, func() {

				if tt.input == "\n" {
					fmt.Println()
					return
				}

				words := strings.Split(tt.input, "\n")
			
				ap.PrintAscii(banner, words)
			})

			if out != tt.expected {
				t.Fatalf("output mismatch\nEXPECTED:\n%s\nGOT:\n%s", tt.expected, out)
			}
		})
	}
}