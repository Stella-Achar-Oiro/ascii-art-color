package Ascii

import (
	"bufio"
	"fmt"
	"os"
)

// GetLine is a function to read font from files
func GetLine(num int, filename string) string {
	// Checks if the banner file size is not altered.
	filePath, err := FileCheck(filename)
	if err != nil {
		fmt.Println(err, filePath)
		os.Exit(1)
	}

	f, e := os.Open("../bannerFiles/" + filename + ".txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line

}
