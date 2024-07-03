package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	// something like error handling if .. the user doesn't provide an argument
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <input>")
	// 	return
	// }
	// input := strings.Join(os.Args[1:], " ")
	// fmt.Println(input)
	ls_cmd := exec.Command("ls")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	// echo "something elsa e" >> smh.txt
}
func create_file() string {
	exec.Command("touch", arg)
}
func add_txt(text) {

}

func create_file(dir string, filename string) error {
	f, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
