package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "aaaaaaa"
	fmt.Println(isAnagram(s, t))
}
