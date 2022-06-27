package main

import (
	"fmt"
	"math"
)

func powerSetIterative(s []int, out *[][]int) {
	size := int(math.Pow(2, float64(len(s))))

	for counter := 0; counter < size; counter++ {
		for j := 0; j < len(s); j++ {
			if counter&(1<<j) > 0 {
				fmt.Println(s[j])
			}
		}
	}
}

func powerSet(s []int, index int, curr []int, out *[][]int) {
	n := len(s)

	if index == n {
		return
	}

	if len(curr) > 0 {
		fmt.Println("output", curr)
		c := make([]int, len(curr))
		copy(c, curr)
		*out = append(*out, c)
	}

	//adding the rest of characters
	for i := index + 1; i < n; i++ {
		fmt.Println("add ", s[i], i)
		curr = append(curr, s[i])
		powerSet(s, i, curr, out)
		fmt.Println("before drop", curr, i)

		//
		curr = curr[:len(curr)-1]
		fmt.Println("after drop", curr, i)
	}
}

func main() {
	s := []int{1, 2, 3}
	var res [][]int
	powerSet(s, -1, nil, &res)
	//powerSetIterative(s, &res)
	fmt.Println(res)
}
