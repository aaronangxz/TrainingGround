package main

import (
	"fmt"

	"github.com/aaronangxz/TrainingGround/common"
)

func RunBinarySearch() {
	a := common.MakeSortedSlice(10)
	fmt.Println("Binary Search Recursive")
	fmt.Println(BinaryRecursive(common.MakeSortedSlice(10), 0, len(a)-1, a[8]))
	fmt.Println("Binary Search Iterative")
	fmt.Println(BinaryIter(common.MakeSortedSlice(10), 0, len(a)-1, a[8]))
}

func BinaryRecursive(slice []int, l int, r int, target int) int {
	if l <= r {
		mid := l + (r-l)/2

		if slice[mid] == target {
			return mid
		}

		if target > slice[mid] {
			return BinaryRecursive(slice, mid+1, r, target)
		}
		return BinaryRecursive(slice, l, mid-1, target)
	}
	return -1
}

func BinaryIter(slice []int, l int, r int, target int) int {
	for l <= r {
		mid := l + (r-l)/2

		if slice[mid] == target {
			return mid
		}

		if target > slice[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	RunBinarySearch()
}
