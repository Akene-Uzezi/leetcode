package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	r := reverseString(s)
	return s == r
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	val := isPalindrome("M o m")
	fmt.Println(val)
}
