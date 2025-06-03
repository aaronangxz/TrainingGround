package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	canOpen := make([]bool, n)
	for i := 0; i < n; i++ {
		canOpen[i] = status[i] == 1
	}

	hasBox := make([]bool, n)
	used := make([]bool, n)
	var q []int
	ans := 0

	for _, box := range initialBoxes {
		hasBox[box] = true
		if canOpen[box] {
			q = append(q, box)
			used[box] = true
			ans += candies[box]
		}
	}

	for len(q) > 0 {
		bigBox := q[0]
		q = q[1:]
		for _, key := range keys[bigBox] {
			canOpen[key] = true
			if !used[key] && hasBox[key] {
				q = append(q, key)
				used[key] = true
				ans += candies[key]
			}
		}
		for _, innerBox := range containedBoxes[bigBox] {
			hasBox[innerBox] = true
			if !used[innerBox] && canOpen[innerBox] {
				q = append(q, innerBox)
				used[innerBox] = true
				ans += candies[innerBox]
			}
		}
	}

	return ans
}
