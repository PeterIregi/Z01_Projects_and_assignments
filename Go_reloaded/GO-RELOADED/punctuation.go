package reloaded


import (
	"strings"

)

func punctuation(text string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	// 1. Remove space BEFORE punctuation
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p)
	}

	// 2. Add space AFTER punctuation
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, p, p+" ")
	}

	// 3. Grouped punctuation
	groups := []string{"...", "!!", "??", "!?", "?!"}
	for _, g := range groups {
		text = strings.ReplaceAll(text, strings.Join(strings.Split(g, ""), " "), g)
	}

	// 4. Quotes: no spaces inside
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	// 5. Normalize spaces
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}

	// 6. Attach punctuation to closing quote
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, p+" '", p+"'")
	}

	// 7. SPECIAL CASE: colon keeps space before opening quote
	text = strings.ReplaceAll(text, ":'", ": '")

	return strings.TrimSpace(text)
}
