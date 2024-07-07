package main

import (
	"fmt"
	"gem/handler"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Help:	go run gem.go --help")
		return
	} else {
		switch os.Args[1] {
		case "-ls", "-rm", "-add":
			text := os.Args[2:]
			if len(text) > 1 {
				pug := strings.Join(text, " ")
				handleCommand(os.Args[1], pug)
			} else {
				pug := strings.Join(text, " ")
				handleCommand(os.Args[1], pug)
			}
		case "--help":
			handler.Help()
		default:
			handler.Help()
		}
	}
}
func handleCommand(cmd string, args string) {
	switch cmd {
	case "-rm":
		handler.Remove_from_JSON(handler.MAIN_FILE, args)
	case "-add":
		add_cmd(args)
	case "-ls":
		handler.Display_JSON(handler.MAIN_FILE, args)
	}
}

func add_cmd(args string) {
	day := time.Now().Weekday().String()
	new_idea := handler.Idea_struct{Idea: args, Day: day}
	handler.Write_JSON(handler.MAIN_FILE, new_idea)
	fmt.Println("<add> Done!")
}
