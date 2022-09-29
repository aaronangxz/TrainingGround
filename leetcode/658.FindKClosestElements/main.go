package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func findClosestElements(arr []int, k int, x int) []int {
	for len(arr) > k {
		if common.Abs(arr[0]-x) > common.Abs(arr[len(arr)-1]-x) {
			arr = arr[1:]
		} else {
			arr = arr[:len(arr)-1]
		}
	}
	return arr
}

func findClosestElementsBinarySearch(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	for left < right {
		mid := (left + right) / 2

		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElementsBinarySearch(a, 4, 3))
	fmt.Println(findClosestElements(a, 4, 3))
}
