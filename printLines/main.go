package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("expecting: go run . [string]")
	// 	os.Exit(0)
	// }
	// input := os.Args[1]
	// if input == " " {
	// 	os.Exit(0)
	// }
	// input = strings.ReplaceAll(input, `\n`, "\n")
	// if input == "\n" {
	// 	fmt.Print("\n")
	// 	os.Exit(0)
	// }
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("failed to read standard file", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if i >= 20 {
			break
		}
		fmt.Printf("line %d: %q\n", i, line)
	}
}
