package main

import (
	"fmt"
	"log"

	"github.com/aaronangxz/TrainingGround/common"
)

func RunSelectionSort() {
	SelectionSort(common.MakeSlice(20))
}

func SelectionSort(slice []int) {
	fmt.Println("Selection Sort")
	fmt.Println("Before:", slice)
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		if err := common.SwapAny(&slice[i], &slice[min]); err != nil {
			log.Fatal(err.Error())
		}
	}
	fmt.Println("After:", slice)
}

func main() {
	RunSelectionSort()
}
