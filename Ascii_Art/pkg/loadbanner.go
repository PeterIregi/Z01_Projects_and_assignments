package pkg

import (
	"os"
	"fmt"
	"strings"
	"log"
)


func LoadBanner(filename string) map[rune][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("file could not be opened")
		return nil
	}

	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if info.Size() < 10 {
		return nil
	}


	lines := strings.Split(string(data), "\n")

	bannerMap := make(map[rune][]string)

	for ch := 32 ; ch <= 126; ch++ {
		start := int(ch -32) * 9
		bannerMap[rune(ch)] = lines[start + 1 : start + 9 ]
	}

	return bannerMap

}