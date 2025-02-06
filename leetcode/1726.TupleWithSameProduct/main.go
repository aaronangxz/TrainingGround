package main

func tupleSameProduct(nums []int) int {
	sumMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sumMap[nums[i]*nums[j]]++
		}
	}
	combinations := 0
	for _, m := range sumMap {
		combinations += (m - 1) * m * 4
	}
	return combinations
}

// 2 x 3 = 6
// 2 x 4 = 8
// 2 x 6 = 12

// 3 x 4 = 12
// 3 x 6 = 18

// 4 x 6 = 24

// [2,6],[3,4]
// Permute
// [2,6,3,4] [6,2,3,4] [2,6,4,3] [6,2,4,3]
// Reverse Permute
