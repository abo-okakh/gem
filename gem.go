package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const MAIN_FILE string = "gem_storage.json"

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
			break
		case "--help":
			help()
		default:
			help()
		}
	}
}

// TODO:[the items will be numbered]
// TODO: command -ls = list of items with times
// TODO: command -rm[arg init] = remove and an item with by putting the number for it
type Idea_struct struct {
	Idea string `json:"idea"`
	Day  string `json:"day"`
}

func Write_JSON(file string, newIdea Idea_struct) {
	// Read the existing file data
	file_data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("File reading error in Write_JSON:", err)
		return
	}

	var ideas []Idea_struct
	// Unmarshal the file data if not empty
	if len(file_data) > 0 {
		err = json.Unmarshal(file_data, &ideas)
		if err != nil {
			fmt.Println("Unmarshal error:", err)
			return
		}
	}
	// 	json.Unmarshal takes JSON-encoded data (file_data) and stores the result in the value pointed to by &ideas.
	// &ideas is a pointer to the slice ideas, allowing Unmarshal to modify the slice directly.

	// Append the new idea
	ideas = append(ideas, newIdea)

	// Marshal the updated slice back to JSON
	updatedJSON, err := json.MarshalIndent(ideas, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the updated JSON back to the file
	err = os.WriteFile(file, updatedJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func handleCommand(cmd string, args string) {
	switch cmd {
	case "-rm":
		fmt.Println("remove")
	case "-add":
		add_cmd(args)
	case "-ls":
		ls_cmd(args)
	}
}

func add_cmd(args string) {
	day := time.Now().Weekday().String()
	new_idea := Idea_struct{Idea: args, Day: day}
	Write_JSON(MAIN_FILE, new_idea)
	fmt.Println("<add> Done!")
}

func ls_cmd(args string) {
	//TODO: my plan this should be sorted by numbers [ my other plan was to make it show from the news to the oldest but not now]
	fmt.Println(args)
}
func help() {
	fmt.Println("-add / to add ideas to gem storage")
	fmt.Println("	eg. go run gem.go -add <your idea>")
	fmt.Println("	")
	fmt.Println("-ls / list the ideas from gem storage")
	fmt.Println("	eg. go run gem.go -ls")
	fmt.Println("	")
	fmt.Println("-remove / remove an idea from gem storage")
	fmt.Println("	eg. go run gem.go -remove[<number of idea from the list using -ls>]")
}
