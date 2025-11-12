package main

func minOperations(nums []int) int {
	n := len(nums)
	ones := 0
	for _, x := range nums {
		if x == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	minLen := n + 1
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}

	if minLen == n+1 {
		return -1
	}
	return (minLen - 1) + (n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
