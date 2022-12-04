package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)

	for _, a := range arr {
		m[a]++
	}

	c := make(map[int]int)

	for _, v := range m {
		c[v]++

		if c[v] > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
