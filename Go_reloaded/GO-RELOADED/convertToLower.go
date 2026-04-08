package reloaded

import (
	"strconv"
	"strings"
)

func convertToLower(text string) string {
	words := strings.Fields(text)

	for i := 0; i< len(words); i++{
		n := 0 
		//deal with instance of "(low)" alone
		if strings.Trim(words[i], "()") == "low"{
			n = 1
		}
		//deal with instance where there is a number after  in the bracket
		if strings.Trim(words[i], "(,") == "low" && i+1 < len(words) && strings.HasSuffix(words[i+1], ")"){
			numStr :=strings.TrimSuffix(words[i+1], ")")
			if num, err := strconv.Atoi(numStr); err ==nil {
				n = num
			}
		}
		//converts word to upper case depending on the value of n
		if n > 0 {
			for j := i - 1; j >= 0 && j >= i - n; j-- {
				words[j] = strings.ToLower(words[j])
			}
			//remove "(low)"
			if words[i] == "(low)" {
				words = append(words[:i], words[i+1:]...)
			}else {
				//remove "(low" and "n)"
				words = append(words[:i], words[i+2:]...)
			}
		}
	}

	return strings.Join(words, " ")
}
