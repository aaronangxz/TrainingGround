package main

func numTilePossibilities(tiles string) int {
	countMap := make(map[string]int)
	for _, t := range tiles {
		countMap[string(t)]++
	}

	var backtrack func() int
	backtrack = func() int {
		res := 0
		for k, v := range countMap {
			if v > 0 {
				countMap[k]--
				res++
				res += backtrack()
				countMap[k]++
			}
		}
		return res
	}
	return backtrack()
}
