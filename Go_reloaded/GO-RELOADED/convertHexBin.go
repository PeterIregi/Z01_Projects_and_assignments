package reloaded

import (
	"strconv"
	"strings"
)

func convertHexBin(text string) string{
	//split string to slice of string with space as separator
	words := strings.Fields(text)

	for i := 1; i < len(words); i++{
		//find if word is (Hex)
		if words[i] == "(hex)"{
			//convert the previous word from hex to decimal
			n, _ :=strconv.ParseInt(words[i-1],16 ,64)
			words[i-1] = strconv.FormatInt(n,10)
			//replace new formated string and remove (hex)
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(bin)"{
			//convert the previous word from binary to decimal
			n, _ :=strconv.ParseInt(words[i-1],2 ,64)
			words[i-1] = strconv.FormatInt(n,10)
			//replace new formated string and remove (bin)
			words = append(words[:i], words[i+1:]...)
			i--
		} 
	}
	return strings.Join(words, " ")
}