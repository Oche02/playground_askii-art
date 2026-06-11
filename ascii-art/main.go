package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ATTENTION: expecting >> go run . [string]")
		return
	}
	input := os.Args[1]

	if input == "data.txt" {
		fmt.Println("<<< READING FROM TEXT FILE >>>")
		fmt.Print("\n")

		read, err := os.ReadFile("data.txt")
		if err != nil {
			fmt.Println("failed to read data.txt file", err)
			os.Exit(1)
		}
		turn := string(read)
		turn = strings.ToLower(turn)
		output := map[string]int{}

		for _, char := range turn {
			output[string(char)]++
		}
		for key, value := range output {
			fmt.Println(key, ">>", value)
		}
		if turn == "" {
			fmt.Println("data.txt file is empty")
			return
		}
	} else {

		fmt.Println("<<< READING FROM TERMINAL >>>")
		input = strings.ToLower(input)

		charMap := map[string]int{}

		for _, text := range input {
			charMap[string(text)]++
		}
		for key, value := range charMap {
			fmt.Println(key, ">>", value)
		}

	}
}
