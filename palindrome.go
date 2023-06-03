package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s) // Convert string to lowercase for case-insensitive comparison
	s = strings.ReplaceAll(s, " ", "") // Remove spaces from the string

	// Compare characters from both ends of the string
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	str := "Able was I saw Elba" // Example palindrome string
	if isPalindrome(str) {
		fmt.Println("It is a palindrome.")
	} else {
		fmt.Println("It is not a palindrome.")
	}
}

