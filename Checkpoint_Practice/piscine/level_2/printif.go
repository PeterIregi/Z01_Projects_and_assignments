package level_2

func PrintIf(str string) string {
	result := "G\n"
	if len(str) < 3 && len(str) > 0 {
		result = "Invalid Input\n"
	}
	return result
}