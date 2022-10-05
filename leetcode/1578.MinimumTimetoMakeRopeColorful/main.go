package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func minCost(colors string, neededTime []int) int {
	res, maxCost, n := 0, 0, len(colors)

	//iterate through string
	for i := 0; i < n; i++ {
		//from index 1 onwards, compare to the previous index
		if i > 0 && colors[i] != colors[i-1] {
			//if they are different, reset the max cost
			maxCost = 0
		}

		//only take the minimum cost between the current max and the current time needed
		res += common.Min(maxCost, neededTime[i])
		//keep the maximum cost rolling
		maxCost = common.Max(maxCost, neededTime[i])
	}
	return res
}

func main() {
	s := "abaac"
	needed := []int{1, 2, 3, 4, 5}
	fmt.Println(minCost(s, needed))
}
