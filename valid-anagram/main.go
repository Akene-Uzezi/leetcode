package main

import "fmt"

func isAnagram(s string, t string) bool {
	sHash := make(map[rune]bool)
	tHash := make(map[rune]bool)
	for _, l := range s {
		sHash[l] = true
	}
	for _, l := range t {
		tHash[l] = true
	}
	return len(sHash) == len(tHash)
}

func main() {
	s := "anagram"
	t := "test"
	val := isAnagram(s, t)
	fmt.Println(val)
}
