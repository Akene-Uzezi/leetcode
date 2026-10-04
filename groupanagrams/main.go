package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	return nil
}

func main() {
	val := groupAnagrams([]string{""})
	fmt.Println(val)
}
