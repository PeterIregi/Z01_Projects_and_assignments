package reloaded


func Modify(text string) string{
	text=capitalize(text)
	text=convertHexBin(text)
	text=convertToUpper(text)
	text=convertToLower(text)
	text=punctuation(text)
	text=aToAnChange(text)
	return text
}