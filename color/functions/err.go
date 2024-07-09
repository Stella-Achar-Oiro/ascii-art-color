package colour

import (
	"fmt"
	"os"
	"strings"
)

func PrintErr() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
	os.Exit(1)
}

func CheckMultipleFlagInput(input []string) {
	var count int
	for _, arg := range input {
		if strings.HasPrefix(arg, "--color=") {
			count++
		}

	}
	if !strings.HasPrefix(input[0], "--color=") {
		PrintErr()
	}
	if count > 1 {
		PrintErr()
	}
}
