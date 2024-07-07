package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const MAIN_FILE string = "gem_storage.json"

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

func Choose_from_JSON(arg int, file string) {
	// just display data
	file_data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("File reading error in Read_JSON:", err)
		return
	} else {
		var ideas []Idea_struct
		if len(file_data) > 0 {
			err = json.Unmarshal(file_data, &ideas)
			if err != nil {
				fmt.Println("Unmarshal error:", err)
				return
			} else {
				var formatted_string string
				if arg >= 0 && arg < len(ideas) {
					formatted_string = fmt.Sprintf("%s %s", ideas[arg].Day, ideas[arg].Idea)
					fmt.Println(formatted_string)
				} else {
					fmt.Println("Invalid index")
				}
			}
		}
	}
}
func Display_JSON(file string, amount interface{}) {
	// just display data
	file_data, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("File reading error in Read_JSON:", err)
		return
	} else {
		var ideas []Idea_struct
		if len(file_data) > 0 {
			err = json.Unmarshal(file_data, &ideas)
			if err != nil {
				fmt.Println("Unmarshal error:", err)
				return
			} else {
				switch pug := amount.(type) {
				case string:
					if pug == "-all" {
						for i := 0; i < len(ideas); i++ {
							fmt.Println(i, "-", "(", ideas[i].Day, ")", ":", ideas[i].Idea)
						}
					} else {
						i, err := strconv.Atoi(pug)
						if err != nil {
							Help()
						} else {
							Choose_from_JSON(i, MAIN_FILE)
						}
					}
				}
			}
		}
	}
}

func Help() {
	fmt.Println("-add / to add ideas to gem storage")
	fmt.Println("	eg. go run gem.go -add <your idea>")
	fmt.Println("	")
	fmt.Println("-ls / list the ideas from gem storage")
	fmt.Println("-ls <number of item>")
	fmt.Println("	eg. go run gem.go -ls")
	fmt.Println("	")
	fmt.Println("-rm / remove an idea from gem storage")
	fmt.Println("	eg. go run gem.go -rm [<number of idea from the list using -ls>]")
}

func Remove_from_JSON(file string, args string) {
	rm_index, err := strconv.Atoi(args)
	if err != nil {
		Help()
	} else {
		file_data, err := os.ReadFile(file)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("error in Removed_from_JSON", err)
			return
		}

		var ideas []Idea_struct
		if len(file_data) > 0 {
			err = json.Unmarshal(file_data, &ideas)
			if err != nil {
				fmt.Println("problem unmarshalling")
				return
			}
		}
		fmt.Println("Removed : ", ideas[rm_index])
		// Removing logic goes here
		ideas = append(ideas[:rm_index], ideas[rm_index+1:]...)
		// updating json
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
}
