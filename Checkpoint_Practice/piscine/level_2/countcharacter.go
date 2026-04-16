package level_2

func CountChar(txt string, c rune) int {
	result := 0
	if len(txt) == 0 {
		result = 0
	} else {
		for _, ch := range txt {
			if ch == c {
				result += 1
			} else {
				continue
			}
		}
	}
	return result
}
