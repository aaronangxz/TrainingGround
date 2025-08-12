package main

func numberOfWays(n int, x int) int {
	cache := make(map[int]int)
	result := count(n, x, 0, 1, cache)
	return result
}

func count(n int, x int, s int, num int, cache map[int]int) int {
	key := (s << 9) + num
	value, exists := cache[key]
	if exists {
		return value
	}
	if s > n {
		return 0
	}
	if s == n {
		return 1
	}
	power := 1
	for i := 0; i < x; i++ {
		power *= num
	}
	if power > n {
		return 0
	}
	modulo := int64(1000000007)
	result := int((int64(count(n, x, s+power, num+1, cache)) + int64(count(n, x, s, num+1, cache))) % modulo)
	cache[key] = result
	return result
}
