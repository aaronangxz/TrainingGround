package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"math/rand"
)

func median(a, aIdx, b, bIdx, c, cIdx int) (int, int) {
	if a < b {
		switch {
		case b < c:
			return b, bIdx
		case a < c:
			return c, cIdx
		default:
			return a, aIdx
		}
	}
	switch {
	case a < c:
		return a, aIdx
	case b < c:
		return c, cIdx
	default:
		return b, bIdx
	}
}

func getPivot(v []int) (int, int) {
	n := len(v)
	r1 := rand.Intn(n)
	r2 := rand.Intn(n)
	r3 := rand.Intn(n)
	return median(v[r1], r1, v[r2], r2, v[r3], r3)
}

func partition(s []int, low, high int) ([]int, int) {
	pivot, pIdx := getPivot(s)

	s[high], s[pIdx] = s[pIdx], s[high]

	i := low

	for j := low; j < high; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[pIdx] = s[pIdx], s[i]

	return s, i
}

func quickSort(s []int, low, high int) []int {
	if low < high {
		var partitionIdx int
		s, partitionIdx = partition(s, low, high)
		s = quickSort(s, low, partitionIdx-1)
		s = quickSort(s, partitionIdx+1, high)
	}
	return s
}

func quickSortStart(s []int) []int {
	low := 0
	high := len(s) - 1
	return quickSort(s, low, high)
}

func quickSortV2(s []int) []int {
	if len(s) < 2 {
		return s
	}

	low, high := 0, len(s)-1

	//get pivot: median of 3 random elements
	pivot, pIdx := getPivot(s)

	//swap pivot to the last
	s[pIdx], s[high] = s[high], s[pIdx]

	//move elements smaller than pivot to the front
	for i := range s {
		if s[i] < pivot {
			s[low], s[i] = s[i], s[low]
			low++
		}
	}

	//no smaller elements, swap pivot back
	//this is the sorted position
	s[low], s[high] = s[high], s[low]

	//sort left and right partition, excluding sorted position
	quickSortV2(s[:low])
	quickSortV2(s[low+1:])

	return s
}

func main() {
	s := common.MakeSlice(50)
	fmt.Println(s)
	//fmt.Println(quickSortStart(s))
	fmt.Println(quickSortV2(s))
}
