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
	testarr := []int{3, 2, 4}
	testtarget := 6
	val := twoSum(testarr, testtarget)
	fmt.Println(val)
}
