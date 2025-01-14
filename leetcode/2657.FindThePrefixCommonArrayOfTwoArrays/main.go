package main

func findThePrefixCommonArray(A []int, B []int) []int {
	// Keeping track of char occurrences
	charMap := make(map[int]int)
	res := make([]int, len(A))

	// Keep track of total count of char appearances
	count := 0

	for i := range A {
		charMap[A[i]]++
		// As long as it appears twice, it is a match on both sides
		// This is possible because A and B are permutations
		if charMap[A[i]] >= 2 {
			count++
		}

		charMap[B[i]]++
		if charMap[B[i]] >= 2 {
			count++
		}
		// Update the result with count of this idx
		res[i] = count
	}
	return res
}

// 1:1 3:1 = [0,0,0,0]
// 3:2 1:2 ++2 = [0,2,0,0]
// 2:1 2:2 ++2++1 = [0,2,3,0]
// 4:1 4:2 ++2++1++1 = [0,2,3,4]
