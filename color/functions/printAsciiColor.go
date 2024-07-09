package colour

import (
	Ascii "asci-art/banner"
	"fmt"
	"os"
	"strings"
)

// PrintAsciiColor prints a word in ASCII art with optional coloring for a specific substring.
func PrintAsciiColor(word string, toColor map[rune]string, wordToColor, fileName string) {

	var r, g, b int
	var err error
	indices, status := findSubstringsInWord(word, wordToColor)

	for i := 0; i < 8; i++ {

		for j, letter := range word {
			// Get the ASCII art line for the current letter.
			line := Ascii.GetLine(1+int(letter-' ')*9+i, fileName)
			// Check if the word needs to be colored and if the current character falls within the start and end indices.
			if status == "subString" {
				if color, found := toColor[letter]; found {
					r, g, b, err = RgbExtract(color)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				if indices[j] == letter {
					fmt.Print(ESCseq(r, g, b), line)
				} else {
					fmt.Print(ESCseq(255, 255, 255), line)
				}

				//checks if the current word is equal to the word that is supposed to be colored and if oK is false
			} else if status == "contains" && strings.ContainsRune(wordToColor, letter) {
				if color, found := toColor[letter]; found {
					r, g, b, err = RgbExtract(color)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
				fmt.Print(ESCseq(r, g, b), line)
			} else {
				// Print the character in white if no specific color is assigned.
				fmt.Print(ESCseq(255, 255, 255), line)
			}

		}
		fmt.Println()
	}

}

func findSubstringsInWord(word, substring string) (map[int]rune, string) {
	track := make(map[int]rune)
	subLen := len(substring)
	wordLen := len(word)

	for i := 0; i <= wordLen-subLen; i++ {
		if word[i:i+subLen] == substring {
			for j := 0; j < subLen; j++ {
				track[i+j] = rune(word[i+j])
			}
		}
	}
	if len(track) == 0 {
		for _, ch := range substring {
			if strings.ContainsRune(word, ch) {
				return nil, "contains"
			}

		}
	}
	if len(track) == 0 {
		fmt.Printf("%s was not found in the text below\n", substring)
	}

	return track, "subString"
}
