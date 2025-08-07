package main

// A solution to #3363 "Find the Maximum Number of Fruits Collected"
//
// As the hints to the problem underline, the core of a solution to the problem
// would be an application of dynamic programming approach. Assuming dp[r][c]
// as the maximum amount of fruits collected by the second or third child, the
// following formulas are required to calculate the array:
//
// a)  the case of the second child, who starts at the right top corner:
//
//	dp[0][n-1] = fruits[0][n-1];
//	dp[r][c] = max(dp[r-1, c-1], dp[r-1, c], dp[r-1, c+1]) + fruits[r][c]
//	for all r from 1 to n-2 and, for each r fixed and set, for all c from
//	max(r+1, n-r-1) to n-2;
//
// b)  the case of the third child, who starts at the left bottom corner:
//
//	dp[n-1][0] = fruits[n-1][0];
//	dp[r][c] = max(dp[r-1, c-1], dp[r, c-1], dp[r+1, c-1]) + fruits[r][c]
//	for all c from 1 to n-2, and, for each c fixed and set, for all r from
//	max(c+1, n-c-1) to n-2.
//
// After the calculations, the resulting maximum amount of fruits would be in
// dp[n-2][n-1] and dp[n-1][n-2] correspondingly. Of course, corner cases of
// indices going out the border of zero and n-1 should be carefully accounted
// for.
//
// This solution has the following features regarding to the approach.
//
//  1. It sports a collapse of 2D dp array to 1D. Indeed, one would need just
//     previous line to calculate the current one. More than that, one would
//     not need the whole previous line. As both max-functions in the formulas
//     above operate over just three elements, one would need to save just two
//     numbers, constantly shifting them.
//
//  2. It uses additional cells with zero values in the dp array to avoid
//     processing of corner cases. That has an advantage of cutting conditions
//     off and it simplifies the resulting code, at least in my (very humble)
//     opinion. Of course, it has a disdavantage, as one would need additional
//     memory. Without the trick one could realize the algorithm completely
//     in-place with O(1) additional memory.
func maxCollectedFruits(fruits [][]int) int {
	n := int16(len(fruits))
	// Note additional cells. The whole indexation will be shifted by 1
	// below, and the first and last elements will always be zero.
	dp := make([]int32, n/2+2)

	dp[1] = int32(fruits[0][n-1])
	for r := int16(1); r < n-1; r++ {
		// Additional variables to roll over previous values during the
		// calculation above. The role of prev0 is to hold a value of
		// previous element _before_ its recalculation. Similarly,
		// prev1 holds a value of current element before the
		// recalculation. Their values are used for the max-functions
		// above instead of dp[r-1, c-1], dp[r-1, c] values (and
		// dp[r-1,c-1], dp[r, c-1] values correspondingly in the code
		// for the third child).
		prev0, prev1 := dp[0], dp[1]
		for c := n - 1; max(r+1, n-r-1) <= c; c-- {
			// It should be mentioned that there is a trick: the
			// values are saved during the iteration _before_ the
			// corresponding recalculation. It can be easily
			// tracked though:
			// before the cycle:    prev0 <~ dp[0], prev1 <~ dp[1];
			// during dp[1] recalc: prev0 <~ dp[1], prev1 <~ dp[2];
			// during dp[2] recalc: prev0 <~ dp[2], prev2 <~ dp[3];
			// and so on
			prev0, prev1, dp[n-c] = prev1, dp[n-c+1], max(prev0, prev1, dp[n-c+1])+int32(fruits[r][c])
		}
	}
	result := dp[1]

	clear(dp)
	dp[1] = int32(fruits[n-1][0])
	for c := int16(1); c < n-1; c++ {
		prev0, prev1 := dp[0], dp[1]
		for r := n - 1; max(c+1, n-c-1) <= r; r-- {
			prev0, prev1, dp[n-r] = prev1, dp[n-r+1], max(prev0, prev1, dp[n-r+1])+int32(fruits[r][c])
		}
	}
	result += dp[1]

	for i := range n {
		result += int32(fruits[i][i])
	}

	return int(result)
}
