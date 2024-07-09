package colour

import (
	"fmt"
	"os"
)

// ColorCheck takes a slice of strings as input and returns a boolean and a slice of strings.
// The boolean indicates if the input contains more than one valid color string.
// The slice of strings contains the valid color strings from the input.
func ColorCheck(c []string) (bool, []string) {
	NewC := make([]string, 0, len(c))
	if len(c) == 0 {
		return false, c
	}

	for _, col := range c {
		if col != "" {
			//checks if rgb can be extracted from each color
			_, _, _, err := RgbExtract(col)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			NewC = append(NewC, col)
		}
	}
	if len(NewC) == 0 || len(NewC) == 1 {
		return false, NewC
	}
	if len(NewC) > len(os.Args[2]) {
		fmt.Println("The colors are more than the words to be colored")
		fmt.Println("Usage: go run. [OPTION] [STRING]  ")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(1)
	}

	return true, NewC

}

// ColorAssighn takes a slice of color strings and assigns them to characters from the second command-line argument.
func ColorAssighn(colors []string) map[rune]string {

	toColor := make(map[rune]string)
	proced, newColors := ColorCheck(colors)

	if proced {
		for i, ch := range os.Args[2] {
			if len(os.Args[2]) >= len(newColors) && i < len(newColors) {
				toColor[ch] = newColors[i]

			} else if i < len(newColors) {
				toColor[ch] = newColors[i]
			} else if i >= len(newColors) {
				if _, ok := toColor[ch]; !ok {
					toColor[ch] = newColors[len(newColors)-1]
				}

			}
		}
		// If there is only one valid color in the input it is assighned to all letters to be colored.
	} else if !proced && len(newColors) == 1 {
		for _, ch := range os.Args[2] {
			toColor[ch] = colors[0]
		}

	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]  ")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(0)
	}
	
	return toColor

}
