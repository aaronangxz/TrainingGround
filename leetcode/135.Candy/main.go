package main

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			c[i] = c[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && c[i] < c[i+1]+1 {
			c[i] = c[i+1] + 1
		}
	}
	sum := 0
	for _, v := range c {
		sum += v
	}
	return sum
}
