package main

import (
	"fmt"

	"github.com/aaronangxz/TrainingGround/common"
)

func insertionSort(s []int) {
	//treat first element as sorted
	for i := 1; i < len(s); i++ {
		curr := s[i]
		j := i - 1

		//Whenever current element is smaller than previous elements,
		//move back one position
		for j >= 0 && curr < s[j] {
			s[j+1] = s[j]
			j--
		}

		//eventually insert into the next sorted position
		s[j+1] = curr
	}
}

func main() {
	slice := common.MakeSlice(20)
	fmt.Println("Before:", slice)
	insertionSort(slice)
	fmt.Println("After:", slice)
}
