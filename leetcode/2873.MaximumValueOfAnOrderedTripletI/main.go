package main

func maximumTripletValue(nums []int) int64 {
	maxTrip := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				val := (nums[i] - nums[j]) * nums[k]
				maxTrip = max(maxTrip, val)
			}
		}
	}
	return int64(maxTrip)
}
