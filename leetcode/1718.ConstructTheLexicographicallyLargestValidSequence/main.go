package main

func constructDistancedSequence(n int) []int {
	size := 2*n - 1
	res := make([]int, size)
	used := make([]bool, n+1)

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		// If we filled the entire sequence, return true
		if idx == size {
			return true
		}
		// If position is already filled, move to the next index
		if res[idx] != 0 {
			return backtrack(idx + 1)
		}

		// Try placing numbers from `n` down to `1`
		for num := n; num >= 1; num-- {
			if used[num] {
				continue
			}
			if num == 1 {
				// Place 1 at index `idx`
				res[idx] = 1
				used[num] = true
				if backtrack(idx + 1) {
					return true
				}
				res[idx] = 0
				used[num] = false
			} else {
				// Place `num` at `idx` and `idx + num`
				if idx+num < size && res[idx+num] == 0 {
					res[idx] = num
					res[idx+num] = num
					used[num] = true
					if backtrack(idx + 1) {
						return true
					}
					// Backtrack
					res[idx] = 0
					res[idx+num] = 0
					used[num] = false
				}
			}
		}
		return false
	}

	backtrack(0)
	return res
}
