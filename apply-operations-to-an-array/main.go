package main

func applyOperations(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}
	write := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[write], nums[i] = nums[i], nums[write]
			write++
		}
	}
	return nums
}

func main() {
	val := applyOperations([]int{1, 2, 2, 1, 1, 0})
	println(val)
}
