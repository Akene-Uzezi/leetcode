package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	val := removeDuplicates([]int{0, 1, 1, 2, 2, 3, 4, 4, 5})
	println(val)
}
