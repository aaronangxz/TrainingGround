package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		//negative because we want to find target + front + back = 0 -> front + back = -target
		target := -nums[i]
		fmt.Println(target)

		//Range to find should be right after the current element, until the end
		front := i + 1
		back := len(nums) - 1

		for front < back {
			sum := nums[front] + nums[back]

			//because it is sorted, so when it is smaller / greater we just move forward / backwards
			if sum < target {
				front++
			} else if sum > target {
				back--
			} else {
				threeSome := []int{nums[i], nums[front], nums[back]}
				res = append(res, threeSome)

				//Skip the next duplicate element, if available
				for front < back && nums[front] == threeSome[1] {
					front++
				}

				//Skip the next duplicate element, if available
				for front < back && nums[back] == threeSome[2] {
					back--
				}
			}
		}

		//Skip the next duplicate element, if available
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
