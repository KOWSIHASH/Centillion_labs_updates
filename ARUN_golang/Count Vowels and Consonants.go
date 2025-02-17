package main

import (
	"fmt"
	"strings"
)

func countVowelsConsonants(s string) (int, int) {
	vowels := "aeiouAEIOU"
	vowelCount, consonantCount := 0, 0

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			consonantCount++
		}
	}
	return vowelCount, consonantCount
}

func main() {
	var text string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	vowels, consonants := countVowelsConsonants(text)
	fmt.Println("Vowels:", vowels, "Consonants:", consonants)
}
