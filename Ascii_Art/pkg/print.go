package pkg

import (
	"fmt"
 	"strings" 
	
)

func PrintAscii(banner map[rune][]string, words []string) {
	
	for _, word := range words {
		if word == ""  || word == " " {
			fmt.Println()
			continue
		}

		word = strings.TrimSpace(word)
	
		for i := 0; i <= 7; i++ {
			for j := 0; j<len(word);j++ {
				char := word[j]
				fmt.Print(banner[rune(char)][i])
			}
			fmt.Println()
		}
	}
}
