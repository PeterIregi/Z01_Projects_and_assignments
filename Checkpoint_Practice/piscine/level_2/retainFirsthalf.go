package level_2

func RetainFirstHalf(str string) string {
	firstHalf := ""
	if len(str)%2 != 0 {
		firstHalf = str[:(len(str)/2)-1]

	} else {
		firstHalf = str[:(len(str)/2)-1]
	}
	return firstHalf
}
