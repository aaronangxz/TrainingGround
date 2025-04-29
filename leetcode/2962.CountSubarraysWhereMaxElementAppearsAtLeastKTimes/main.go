package main

func countSubarrays(nums []int, k int) int64 {
	var total int64
	maxElement := nums[0]
	maxIndices := make([]int, 0, k+1) // Stores indices where maxElement occurs

	for i, val := range nums {
		if val == maxElement {
			// Append index if value equals current maxElement
			maxIndices = append(maxIndices, i)
		} else if val > maxElement {
			// New maxElement found — reset state
			maxElement = val
			maxIndices = append(maxIndices[:0], i) // Clear and insert
			total = 0
		}

		// Keep at most k maxElement indices
		if len(maxIndices) > k {
			maxIndices = maxIndices[1:]
		}

		// If exactly k maxElements in the current window, count valid subarrays
		if len(maxIndices) == k {
			total += int64(maxIndices[0] + 1)
		}
	}

	return total
}
