package main

import (
	"fmt"
	"log"

	"github.com/aaronangxz/TrainingGround/common"
)

func dfs(i int, nums []int, out *[][]int) {
	if i == len(nums) {
		//copy nums into a new slice
		c := make([]int, len(nums))
		copy(c, nums)
		*out = append(*out, c)
		return
	}

	for j := i; j < len(nums); j++ {
		common.Swap(&nums[i], &nums[j])
		log.Println("aft swap:", nums)
		dfs(i+1, nums, out)
		common.Swap(&nums[i], &nums[j])
		log.Println("aft swap:", nums)
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
