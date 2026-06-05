package level_3

func HashCode(dec string) string{
	size := len(dec)
	result:= []rune{}

	for i , ch := range dec{
		h := (int(ch) + size)%127
		if h < 33 {
			h +=33
		}
		result [i]=rune(h)
	}
	return string(result)
}