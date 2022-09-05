package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"math"
)

func nextGreaterElement(n int) int {
	//fill numbers into an itn slice with correct order
	var elements []int
	for n > 0 {
		elements = append(elements, n%10)
		n /= 10
	}
	common.ReverseSliceRange(&elements, 0, len(elements)-1)

	i := len(elements) - 2

	//find the first element that is non-increasing
	//because if it is increasing, there will be no permutation to make it bigger
	//e.g. 5432 -> no greater element
	//e.g. 1234-> greater element is 1243
	for i >= 0 && elements[i] >= elements[i+1] {
		i--
	}

	//we do not found any non-increasing number, hence impossible to find next greater element
	if i == -1 {
		return -1
	}

	j := len(elements) - 1

	//find the next greater number in the increasing portion
	//e.g. 6587654, since i is 5 (firs non-increasing), j will be 6 (first greater than i)
	for j > i && elements[j] <= elements[i] {
		j--
	}

	//swap i and j
	common.Swap(&elements[j], &elements[i])

	//reverse all elements after i
	//the subarray after i should be in decreasing order, so we just have to reverse it inctead of sorting it ASC
	common.ReverseSliceRange(&elements, i+1, len(elements)-1)

	//convert back to int
	res := 0
	for k := 0; k < len(elements); k++ {
		res = res*10 + elements[k]
	}

	//exceeds boundary
	if res > math.MaxInt32 {
		return -1
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement(65876))
	fmt.Println(nextGreaterElement(21))
}
