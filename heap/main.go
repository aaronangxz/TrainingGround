package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"github.com/aaronangxz/TrainingGround/heap/min_heap"
)

func main() {
	h := min_heap.NewMinHeap(11)
	slice := common.MakeSlice(10)
	fmt.Println(slice)

	for _, s := range slice {
		h.InsertKey(s)
	}

	for i := 0; i < len(slice); i++ {
		fmt.Print(h.ExtractMin())
	}

	//fmt.Println(h.ExtractMin())
	//fmt.Println(h.GetMin())
}
