package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	maxVal := 0
	maxLeft := nums[0] // tracks maximum nums[i] to the left of j
	maxDiff := 0       // tracks maximum (nums[i] - nums[j]) for i < j

	for k := 2; k < n; k++ {
		// Update maxDiff with the current j (which is k-1)
		// maxDiff = max(maxDiff, maxLeft - nums[k-1])
		if maxLeft-nums[k-1] > maxDiff {
			maxDiff = maxLeft - nums[k-1]
		}

		// Update maxLeft with the current j (which could be i for future k)
		if nums[k-1] > maxLeft {
			maxLeft = nums[k-1]
		}

		// Calculate current triplet value
		current := maxDiff * nums[k]
		if current > maxVal {
			maxVal = current
		}
	}

	if maxVal < 0 {
		return 0
	}
	return int64(maxVal)
}
