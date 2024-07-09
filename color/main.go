package main

import (
	colour "asci-art/color/functions"
	"fmt"
	"os"
	"strings"

	Ascii "asci-art/banner"
)

func main() {
	color := ""
	str := ""
	fileName := ""

	if len(os.Args) == 2 && !strings.HasPrefix(os.Args[1], "--color=") {
		Ascii.PrintAscii(os.Args[1])
		return
	}
	// Check for the correct number of arguments.
	input := os.Args[1:]
	colour.CheckMultipleFlagInput(input)
	if len(input) < 2 || len(input) > 4 {
		colour.PrintErr()
	}

	if len(input) == 4 && !(input[len(input)-1] == "thinkertoy" || input[len(input)-1] == "shadow" || input[len(input)-1] == "standard") {
		colour.PrintErr()
	}
	fileName = "standard"

	for i, v := range input {

		if len(v) >= 8 && v[0:8] == "--color=" {
			color = v[8:]
		} else if len(v) == 2 && v[0:2] == "--" {
			continue
		} else if (v == "standard" || v == "thinkertoy" || v == "shadow") && i == len(input) {
			fileName = v
		} else {
			if len(input) == 4 {
				str = input[2]
			} else if len(input) == 3 {
				str = input[2]
			} else if len(input) == 2 || (len(input) == 3 && input[len(input)-1] == "standard" || input[len(input)-1] == "thinkertoy" || input[len(input)-1] == "shadow") {
				str = input[1]
			}
		}
	}
	if len(color) == 0 {
		colour.PrintErr()
	}

	str = strings.Replace(str, "\\n", "\n", -1)

	// Finds and displays an error for non-printable characters in the str string if found.
	str = Ascii.HandleSpecialCase(str)

	if str == "\n" {
		fmt.Println()
		return
	} else if str == "" {
		return
	}
	//colors are assighned to each letter to be colored  respectively
	colors := strings.Split(color, "/")
	toColor := colour.ColorAssighn(colors)

	// Split the str into lines based on newline characters.
	Input := strings.Split(str, "\n")
	wordsTobeColored := os.Args[2]
	spaceCount := 0

	// Iterate over each line of the input.
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				fmt.Println()
			}
		} else {
			colour.PrintAsciiColor(word, toColor, wordsTobeColored, fileName)

		}
	}

}
