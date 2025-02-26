package main

func maxAbsoluteSum(nums []int) int {
	// Initialize minPrefixSum and maxPrefixSum to track the minimum and maximum prefix sums.
	minPrefixSum := 0
	maxPrefixSum := 0

	// Initialize prefixSum to accumulate the sum of elements as we iterate through the array.
	prefixSum := 0

	// Iterate through the array to compute the prefix sums and track the min and max prefix sums.
	for _, num := range nums {
		prefixSum += num

		// Update minPrefixSum to the minimum value between the current minPrefixSum and prefixSum.
		minPrefixSum = min(minPrefixSum, prefixSum)

		// Update maxPrefixSum to the maximum value between the current maxPrefixSum and prefixSum.
		maxPrefixSum = max(maxPrefixSum, prefixSum)
	}

	// The maximum absolute sum is the difference between the maximum and minimum prefix sums.
	return maxPrefixSum - minPrefixSum
}
