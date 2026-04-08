package main

import (
	"os"
	"fmt"
	"regexp"
	// "strings"
	ap "ascii_art/pkg"
)
func main(){
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]

	if input == "\n" {
		fmt.Println()
		return
	}

	
	re := regexp.MustCompile(`\\[n]`) // matches \n or \t
	words := re.Split(input, -1)
	

	banner := ap.LoadBanner("./standard.txt")
	ap.PrintAscii(banner,words)
}