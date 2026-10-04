package main

import "fmt"

func twoSum(nums []int, target int) []int {
	incdices := make(map[int]int)
	for i, n := range nums {
		incdices[n] = i
	}
	for i, n := range nums {
		difference := target - n
		if j, found := incdices[difference]; found && j != i {
			return []int{i, j}
		}
	}
	return []int{}
}

func main() {
	testarr := []int{4, 5, 6}
	testtarget := 10
	val := twoSum(testarr, testtarget)
	fmt.Println(val)
}
