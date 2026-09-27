package main

import "fmt"

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		if !isPalindrome(words[i]) {
			continue
		} else {
			return words[i]
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	compare := s == string(runes)
	return compare
}

func main() {
	val := isPalindrome("mom")
	fmt.Println(val)
}
