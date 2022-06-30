package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	count, index int
}

func firstUniqChar(s string) int {
	m := make([]Pair, 26)

	for i := len(s) - 1; i >= 0; i-- {
		m[int(s[i]-'a')] = Pair{m[int(s[i]-'a')].count + 1, i}
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].index < m[j].index
	})

	fmt.Println(m)

	for _, mm := range m {
		if mm.count == 1 {
			return mm.index
		}
	}

	return -1
}

func firstUniqCharNaive(s string) int {
	m := make(map[int]int)

	for _, ss := range s {
		m[int(ss)]++
	}

	for i, ss := range s {
		if m[int(ss)] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
	fmt.Println(firstUniqCharNaive(s))
}
