package asciiarts

import "bufio"
import "os"

func BuildASCII(input []string, charmap map[rune][]string) []string {
	var res []string

	for _, word := range input {
		if word == "" {
			res = append(res, "")
			continue
		}
		for i := 0; i < 8; i++ {
			var line string
			for _, letter := range word {
				ascii, exist := charmap[letter]
				if !exist {
					// Skip unsupported characters or replace
					continue
				}
				line += ascii[i]
			}
			res = append(res, line)
		}
	}
	return res
}

func MapToASCII(lines []string) map[rune][]string {
	charmap := make(map[rune][]string)
	start := 1

	for char := ' '; char <= '~'; char++ {
		charmap[char] = lines[start : start+8]
		start += 9
	}

	return charmap
}

func OutputCharacters(input []string, charmap map[rune][]string) string {
	res := ""
	for _, line := range BuildASCII(input, charmap) {
		res += line + "\n"
	}
	return res
}

func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}