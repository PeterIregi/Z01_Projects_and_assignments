package reloaded

import (
	"strconv"
	"strings"
)

func convertToUpper(text string) string {
	words := strings.Fields(text)

	for i := 0; i< len(words); i++{
		n := 0 
		//deal with instance of "(up)" alone
		if strings.Trim(words[i], "(),") == "up"{
			n = 1
		}
		//deal with instance where there is a number ufter up in the bracket
		if strings.Trim(words[i], "(,") == "up" && i+1 < len(words) && strings.HasSuffix(words[i+1], ")"){
			numStr :=strings.TrimSuffix(words[i+1], ")")
			if num, err := strconv.Atoi(numStr); err ==nil {
				n = num
			}
		}
		//converts word to upper case depending on the value of n
		if n > 0 {
			for j := i - 1; j >= 0 && j >= i - n; j-- {
				words[j] = strings.ToUpper(words[j])
			}
			//remove "(up)"
			if words[i] == "(up)" {
				words = append(words[:i], words[i+1:]...)
			}else {
				//remove "(up" and "n)"
				words = append(words[:i], words[i+2:]...)
			}
		}
	}

	return strings.Join(words, " ")
}
