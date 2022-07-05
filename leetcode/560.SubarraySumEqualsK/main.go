package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	//create prefix sum
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	m := make(map[int]int)

	ans := 0

	for i := 0; i < len(nums); i++ {
		if prefix[i] == k {
			ans++
		}

		if val, found := m[prefix[i]-k]; found {
			ans += val
		}
		m[prefix[i]]++
	}
	return ans
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
