package level_2

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	return str[:len(str)/2]
}
