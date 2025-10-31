package main

func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	seen := make([]bool, n)
	res := make([]int, 0, 2)
	for _, x := range nums {
		if seen[x] {
			res = append(res, x)
			if len(res) == 2 {
				break
			}
		} else {
			seen[x] = true
		}
	}
	return res
}
