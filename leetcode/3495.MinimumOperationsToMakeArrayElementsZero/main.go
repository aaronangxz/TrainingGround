package main

func minOperations(queries [][]int) int64 {
	var ans int64 = 0
	for _, query := range queries {
		start, end := query[0], query[1]
		var ops int64 = 0
		var prev int64 = 1
		for d := int64(1); d < 21; d++ {
			cur := prev * 4
			l := int64(start)
			if prev > l {
				l = prev
			}
			r := int64(end)
			if cur-1 < r {
				r = cur - 1
			}
			if r >= l {
				ops += (r - l + 1) * d
			}
			prev = cur
		}
		ans += (ops + 1) / 2
	}
	return ans
}
