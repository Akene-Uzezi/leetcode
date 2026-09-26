package main

import "fmt"

func moveZeros(nums []int) {
	write := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[write], nums[i] = nums[i], nums[write]
			write++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	fmt.Println(nums)
}
