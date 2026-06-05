package level_3

import (
	"fmt"
	"os"
)

func SearchAndReplace() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	from := os.Args[2]
	to := os.Args[3]

	if len(from) != 1 || len(to) != 1 {
		fmt.Println(str)
		return
	}

	old := from[0]
	new := to[0]

	found := false
	result := []rune{}

	for _, ch := range str {
		if byte(ch) == old {
			result = append(result, rune(new))
			found = true
		} else {
			result = append(result, ch)
		}
	}
	//if nothing as replaced just print original string
	if !found {
		fmt.Println(str)
		return
	}
	fmt.Println(string(result))
}
