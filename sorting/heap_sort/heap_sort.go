package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"log"
)

func heapify(s []int, n, i int) {
	//root is the largest
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	//if left child is larger than root
	if l < n && s[l] > s[largest] {
		largest = l
	}

	//if right child is larger than root
	if r < n && s[r] > s[largest] {
		largest = r
	}

	//if at the end largest is no longer root
	if largest != i {
		if err := common.SwapAny(&s[largest], &s[i]); err != nil {
			log.Fatal(err.Error())
		}
		heapify(s, len(s), largest)
	}
}

func heapSort(s []int) {
	//create max heap starting from one level above the leaf nodes -> n/2 - 1
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapify(s, len(s), i)
	}

	//create max heap again on the reduced array
	for i := len(s) - 1; i > 0; i-- {
		if err := common.SwapAny(&s[0], &s[i]); err != nil {
			log.Fatal(err.Error())
		}
		//the last element is sorted, can ignore for now
		heapify(s[:i], i, 0)
	}
}

func main() {
	s := common.MakeSlice(50)
	fmt.Println(s)
	heapSort(s)
	fmt.Println(s)
}
