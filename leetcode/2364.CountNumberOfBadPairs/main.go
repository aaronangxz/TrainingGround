package main

func countBadPairs(nums []int) int64 {
	// Map to store the frequency of nums[i] - i
	freqMap := make(map[int]int)

	totalBadPairs := 0 // Variable to store total bad pairs

	// Iterate through the array to calculate nums[i] - i
	for i, num := range nums {
		diff := num - i // Calculate nums[i] - i

		// Count of previous positions with same difference (makes a good pair)
		goodPairs := freqMap[diff]

		// Total possible pairs minus good pairs = bad pairs
		totalBadPairs += i - goodPairs

		// Update count of positions with this difference
		freqMap[diff] = goodPairs + 1
	}

	return int64(totalBadPairs)
}
