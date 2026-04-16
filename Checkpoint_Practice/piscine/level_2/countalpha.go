package level_2

func CountAlpha(txt string) int {
	result := 0
	for _, ch := range txt {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			result += 1
		} else {
			continue
		}
	}
	return result
}
