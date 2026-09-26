package main

func removeDuplicates(nums []int) int {
	k := 0
	for i, num := range nums {
		current := num
		next := nums[i+1]
		if current == next {
			k++
		}
	}
	return k
}
