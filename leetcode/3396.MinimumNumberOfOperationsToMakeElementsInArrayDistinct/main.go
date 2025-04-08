package main

func minimumOperations(nums []int) int {
	result := 0
	for len(nums) > 0 {
		if isArrayDistinct(nums) {
			break
		}
		nums = nums[min(3, len(nums)):]
		result++
	}
	return result
}

func isArrayDistinct(nums []int) bool {
	unique := make(map[int]bool)
	for _, num := range nums {
		if unique[num] {
			return false
		}
		unique[num] = true
	}
	return true
}
