package main

func countLargestGroup(n int) int {
	m := make(map[int]int)
	max := 0
	for i := 1; i <= n; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		m[sum]++
		if m[sum] > max {
			max = m[sum]
		}
	}
	count := 0
	for _, v := range m {
		if v == max {
			count++
		}
	}
	return count
}
