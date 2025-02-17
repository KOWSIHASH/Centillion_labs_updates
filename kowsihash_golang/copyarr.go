package main

import "fmt"

func main() {
	src := [5]int{1, 2, 3, 4, 5}
	dest := src
	fmt.Println("Source Array:", src)
	fmt.Println("Copied Array:", dest)
}
