package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func dfs(i int, nums []int, out *[][]int) {
	if i == len(nums) {
		//copy nums into a new slice
		c := make([]int, len(nums))
		copy(c, nums)
		*out = append(*out, c)
	}

	//swap digits, and go one more level down until we exhaust the array elements
	for j := i; j < len(nums); j++ {
		common.Swap(&nums[i], &nums[j])
		dfs(i+1, nums, out)
		common.Swap(&nums[i], &nums[j])
	}
}

func permute(nums []int) [][]int {
	var out [][]int
	dfs(0, nums, &out)
	return out
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
