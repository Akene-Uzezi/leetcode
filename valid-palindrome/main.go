package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
	}
	return false
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
