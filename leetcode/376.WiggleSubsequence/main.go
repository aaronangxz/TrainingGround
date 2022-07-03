package main

import "fmt"

func wiggleMaxLength(nums []int) int {
	//the first element can be both wiggle up or wiggle down
	up, down := 1, 1

	//check from index 1 onwards
	for i := 1; i < len(nums); i++ {
		//if current element is greater than previous, it is wiggling up
		//increment from the previous down because the previous must be wiggling down
		if nums[i] > nums[i-1] {
			up = down + 1
		}
		//if current element is smaller than previous, it is wiggling down
		//increment from the previous up because the previous must be wiggling up
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	//tha maximum is the max wiggle length
	return max(up, down)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength(s))

	s1 := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Println(wiggleMaxLength(s1))

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(wiggleMaxLength(s2))
}
