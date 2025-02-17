package main
import (
	"fmt"
	"os"
)
func writeFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	file.WriteString(content)
	fmt.Println("File written successfully!")
}
func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(data))
}
func main() {
	var filename = "sample.txt"
	var content = "Hello, this is a test file!"
	writeFile(filename, content)
	readFile(filename)
}
