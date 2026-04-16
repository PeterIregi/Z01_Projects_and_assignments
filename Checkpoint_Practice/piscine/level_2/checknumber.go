package level_2

func CheckNumber(txt string) bool {
	result := false
	for _, ch := range txt {
		if ch >= '0' && ch <= '9' {
			result = true
			break

		} else {
			result = false
			continue
		}
	}
	return result
}
