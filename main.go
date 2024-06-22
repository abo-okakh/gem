package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// something like error handling if .. the user doesn't provide an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input>")
		return
	}
	input := strings.Join(os.Args[1:], " ")
	fmt.Println(input)
}
