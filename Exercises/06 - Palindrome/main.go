package main

import (
	"fmt"
	"strings"
)

// Runes supports UNICODE for special characters
func reverseString(text string) string {
	runes := []rune(text)
	length := len(runes)
	for index := 0; index < length/2; index++ {
		runes[index], runes[length-1-index] = runes[length-1-index], runes[index]
	}
	return string(runes)
}

func isPalindrome(message string) bool {
	message = strings.ToLower(message)
	return message == reverseString(message)
}

func main() {
	var word string

	fmt.Println("Please enter a word")
	fmt.Scan(&word)

	if isPalindrome(word) {
		fmt.Println("Is Palindrome")
	} else {
		fmt.Println("Is not Palindrome")
	}
}
