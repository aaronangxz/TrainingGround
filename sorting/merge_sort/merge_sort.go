package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func mergeSort(s []int) []int {
	//when array is sliced until it has only one element, stop slicing
	if len(s) < 2 {
		return s
	}
	mid := len(s) / 2

	//merge the left and right portions
	return merge(mergeSort(s[:mid]), mergeSort(s[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	insertIdx := 0
	i, j := 0, 0

	//two pointers to iterate through both sides
	//insert the smaller element and move its pointer forward
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[insertIdx] = left[i]
			i++
		} else {
			result[insertIdx] = right[j]
			j++
		}
		insertIdx++
	}

	//at least one of the sides is empty now
	//fill the remaining elements to the back of sorted array
	for i < len(left) {
		result[insertIdx] = left[i]
		i++
		insertIdx++
	}

	for j < len(right) {
		result[insertIdx] = right[j]
		j++
		insertIdx++
	}
	return result
}

func main() {
	s := common.MakeSlice(50)
	fmt.Println(s)
	fmt.Println(mergeSort(s))
}
