package main

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)
	maxFreq := 0
	total := 0

	for _, n := range nums {
		freq[n]++
		count := freq[n]

		if count > maxFreq {
			maxFreq = count
			total = count // reset total
		} else if count == maxFreq {
			total += count
		}
	}

	return total
}
