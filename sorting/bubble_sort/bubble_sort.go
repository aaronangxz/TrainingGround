package main

import (
	"fmt"
	"log"

	"github.com/aaronangxz/TrainingGround/common"
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
				if err := common.SwapAny(&slice[j], &slice[j+1]); err != nil {
					log.Fatal(err.Error())
				}
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
			if slice[j] > slice[j+1] {
				if err := common.SwapAny(&slice[j], &slice[j+1]); err != nil {
					log.Fatal(err.Error())
				}
				swapped = true
			}
			if !swapped {
				break
			}
		}
	}
	fmt.Println("After:", slice)
}

func main() {
	RunBubbleSort()
}
