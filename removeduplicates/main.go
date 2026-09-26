package main

func removeDuplicates(nums []int) int {
	k := 0
	var next int
	for i, num := range nums {
		current := num
		if i+1 < len(nums) {
			next = nums[i+1]
		}
		if current == next {
			k++
		}
	}
	return len(nums) - k
}

func main() {
	val := removeDuplicates([]int{1, 1, 2})
	println(val)
}
