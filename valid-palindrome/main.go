package main

import (
	"fmt"
	"unicode"
)

// isPalindrome returns true if s is a palindrome considering only alphanumeric
// characters and ignoring cases.
func isPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		// Skip non-alphanumeric characters from the left
		if !isAlphaNumeric(runes[i]) {
			i++
			continue
		}
		// Skip non-alphanumeric characters from the right
		if !isAlphaNumeric(runes[j]) {
			j--
			continue
		}

		// Case-insensitive comparison
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

// isAlphaNumeric reports whether the rune is a letter or a number.
func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
		{"Was it a car or a cat I saw?", true},
		{"Tab A Cat", false},
		{"Red rum, sir, is murder", true},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		fmt.Printf("Input: %-35q | Got: %-5t | Expected: %t\n", tc.input, result, tc.expected)
	}
}
