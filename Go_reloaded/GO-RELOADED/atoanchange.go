package reloaded

import(
	"strings"
)


func aToAnChange(text string) string{
	text=strings.ReplaceAll(text, "a a", "an a")
	text=strings.ReplaceAll(text, "A a", "An a")

	text=strings.ReplaceAll(text, "a e", "an e")
	text=strings.ReplaceAll(text, "A e", "An e")

	text=strings.ReplaceAll(text, "a i", "an i")
	text=strings.ReplaceAll(text, "A i", "An i")

	text=strings.ReplaceAll(text, "a o", "an o")
	text=strings.ReplaceAll(text, "A o", "An o")

	text=strings.ReplaceAll(text, "a u", "an u")
	text=strings.ReplaceAll(text, "A u", "An u")

	text=strings.ReplaceAll(text, "a h", "an h")
	text=strings.ReplaceAll(text, "A h", "An h")

	return text
}