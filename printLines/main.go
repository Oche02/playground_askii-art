package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error: failed to read standard text file", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if i >= 20 {
			break
		}
		fmt.Printf("Line %d: %q\n", i, line)
	}
}
