package main

import (
	"fmt"
	"sort"
)

type wl struct {
	win  int
	lost int
}

func findWinners(matches [][]int) [][]int {
	winLose := make(map[int]wl)

	for _, match := range matches {
		if v, ok := winLose[match[0]]; ok {
			winLose[match[0]] = wl{v.win + 1, v.lost}
		} else {
			winLose[match[0]] = wl{1, 0}
		}

		if v, ok := winLose[match[1]]; ok {
			winLose[match[1]] = wl{v.win, v.lost + 1}
		} else {
			winLose[match[1]] = wl{0, 1}
		}
	}

	var w, l []int

	for k, v := range winLose {
		if v.win > 0 && v.lost == 0 {
			w = append(w, k)
		} else if v.lost == 1 {
			l = append(l, k)
		}
	}

	sort.Ints(w)
	sort.Ints(l)
	return [][]int{w, l}
}

func main() {
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}
