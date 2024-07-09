package Ascii

import (
	"fmt"
	"strings"
)

func PrintAscii(str string) {

	str = strings.Replace(str, "\\n", "\n", -1)

	str = HandleSpecialCase(str)

	if str == "\n" {
		fmt.Println()
		return
	} else if str == "" {
		return
	}

	spaceCount := 0
	asciiline := ""

	input := strings.Split(str, "\n")
	for _, word := range input {
		if word == "" {
			spaceCount++
			if spaceCount < len(input) {
				fmt.Println()
			}
		} else {

			for i := 0; i < 8; i++ {
				for _, letter := range word {
					asciiline = GetLine(1+int(letter-' ')*9+i, "standard")
					fmt.Print(asciiline)
				}
				fmt.Println()
			}

		}
	}
}
