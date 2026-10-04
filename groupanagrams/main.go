package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	var group [][]string
	for i := 1; i < len(strs); i++ {
		for j := 0; j < i; j++ {
			if isAnagram(strs[i], strs[j]) {
				group = append(group, []string{strs[i], strs[j]})
				continue
			} else {
				group = append(group, []string{strs[j]})
			}
		}
	}
	return group
}

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
	val := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(val)
}
