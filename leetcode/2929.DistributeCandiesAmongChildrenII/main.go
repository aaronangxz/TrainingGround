package main

func distributeCandies(n int, limit int) int64 {
	var res int64 = 0
	m := min(limit, n)
	for i := 0; i <= m; i++ {
		s := n - i
		l := max(0, s-limit)
		r := min(limit, s)
		if r >= l {
			res += int64(r - l + 1)
		}
	}
	return res
}
