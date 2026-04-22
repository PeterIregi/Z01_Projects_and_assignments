package level_2

func RectPerimeter(w, h int) int {
	result := 0
	if w < 0 || h < 0 {
		result = -1
	} else {
		result = 2 * (w + h)
	}

	return result
}