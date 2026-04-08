package reloaded

import (
	"strconv"
	"strings"
)
func transform(word string) string{
	if len(word) == 0{
		return word
	}
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}
func capitalize(text string) string {
	words := strings.Fields(text)

	for i := 0; i< len(words); i++{
		n := 0 
		//deal with instance of "(cap)" alone
		if strings.Trim(words[i], "()") == "cap"{
			n = 1
		}
		//deal with instance where there is a number after  in the bracket
		if strings.Trim(words[i], "(,") == "cap" && i+1 < len(words) && strings.HasSuffix(words[i+1], ")"){
			numStr :=strings.TrimSuffix(words[i+1], ")")
			if num, err := strconv.Atoi(numStr); err ==nil {
				n = num
			}
		}
		//converts word to upper case depending on the value of n
		if n > 0 {
			for j := i - 1; j >= 0 && j >= i - n; j-- {
				words[j] = transform(words[j])
			}
			//remove "(cap)"
			if words[i] == "(cap)" {
				words = append(words[:i], words[i+1:]...)
			}else {
				//remove "(cap" and "n)"
				words = append(words[:i], words[i+2:]...)
			}
			i--
		}
	}

	return strings.Join(words, " ")
}
