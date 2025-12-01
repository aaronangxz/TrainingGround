package main

func maxRunTime(n int, batteries []int) int64 {
	var total int64
	for _, b := range batteries {
		total += int64(b)
	}

	left, right := int64(0), total/int64(n)

	for left < right {
		mid := (left + right + 1) / 2
		need := mid * int64(n)
		var have int64

		for _, b := range batteries {
			if int64(b) < mid {
				have += int64(b)
			} else {
				have += mid
			}
		}

		if have >= need {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}
