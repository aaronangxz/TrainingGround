package main

import "fmt"

func climbStairs(n int) int {
	//whenever n is 2 or less, we only have n ways of climbing
	//n = 1, can only climb 1 step each time for 1 time
	//n = 2, can only climb 1 step each time for 2 times + climb 2 steps each time for 1 time
	if n <= 2 {
		return n
	}

	//n is at least 3 here
	//previous ways are prev / prev2
	prev, prev2, res := 2, 1, 0

	for i := 3; i <= n; i++ {
		//increment all the previous ways
		res = prev + prev2
		//update the previous values
		prev2 = prev
		prev = res
	}
	return res
}
func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
