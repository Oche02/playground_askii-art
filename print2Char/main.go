package main

import (
	"fmt"
	"strings"
)

func main() {
	var result strings.Builder

	for row := 0; row < 8; row++ {
		result.WriteString("some text")
		result.WriteString("\n")
	}

	fmt.Print(result.String())
}
