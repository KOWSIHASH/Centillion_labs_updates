package main
import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("File content:", string(data))
}
