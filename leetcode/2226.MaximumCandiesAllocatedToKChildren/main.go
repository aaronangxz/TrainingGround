package main

func maximumCandies(candies []int, k int64) int {
	// Calculate total candies
	var totalCandies int64 = 0
	for _, c := range candies {
		totalCandies += int64(c)
	}

	// If candies are not even enough for k, it is not possible
	if totalCandies < k {
		return 0
	}

	// Start binary search from 1 - total/k (theoretical max)
	l := int64(1)
	r := totalCandies / k
	res := 0

	for l <= r {
		mid := l + (r-l)/2

		var total int64 = 0
		for _, c := range candies {
			// Pile can be split into n smaller piles
			if int64(c) >= mid {
				total += int64(c) / mid
			}
			// Early exit if it is already more than enough
			if total >= k {
				break
			}
		}

		// Found an answer so far, move to a bigger possible pile
		if total >= k {
			res = int(mid)
			l = mid + 1
		} else {
			// Othersie, downsize the pile
			r = mid - 1
		}
	}

	return res
}
