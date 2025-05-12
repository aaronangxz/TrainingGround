package main

func findEvenNumbers(digits []int) []int {
	// Count the frequency of each digit in the input
	var available [10]int
	for _, digit := range digits {
		available[digit]++
	}

	var result []int

	// Try all 3-digit even numbers (from 100 to 998, step 2)
tryNextNumber:
	for num := 100; num < 1000; num += 2 {
		digitFreq := countDigits(num)

		// Check if this number can be formed from the available digits
		for digit := 0; digit < 10; digit++ {
			if digitFreq[digit] > available[digit] {
				continue tryNextNumber
			}
		}

		result = append(result, num)
	}

	return result
}

// Helper: count how many times each digit appears in a number
func countDigits(num int) [10]int {
	var freq [10]int
	for num > 0 {
		digit := num % 10
		freq[digit]++
		num /= 10
	}
	return freq
}
