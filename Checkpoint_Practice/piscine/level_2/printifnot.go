package level_2

func PrintIfNot(str string) string {
	result := "G\n"
	if len(str) >= 3 {
		result = "Invalid Input\n"
	}
	return result
}