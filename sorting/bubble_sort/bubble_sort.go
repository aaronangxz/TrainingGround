package main

import (
	"fmt"

	"github.com/aaronangxz/XZTrainingGround/common"
)

func RunBubbleSort() {
	BubbleSort(common.MakeSlice(20))
	BubbleSortOptimized(common.MakeSlice(20))
}

func BubbleSort(slice []int) {
	fmt.Println("Bubble Sort")
	fmt.Println("Before:", slice)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				common.Swap(&slice[j], &slice[j+1])
			}
		}
	}
	fmt.Println("After:", slice)
}

func BubbleSortOptimized(slice []int) {
	fmt.Println("Bubble Sort Optimized")
	fmt.Println("Before:", slice)
	swapped := false
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			//log.Printf("Checking %v %v\n", slice[j], slice[j+1])
			if slice[j] > slice[j+1] {
				//log.Printf("Swapping %v %v\n", slice[j], slice[j+1])
				common.Swap(&slice[j], &slice[j+1])
				swapped = true
			}
			if !swapped {
				break
			}
		}
		//log.Printf("Now: %v\n", slice)
	}
	fmt.Println("After:", slice)
}

func main() {
	RunBubbleSort()
}
