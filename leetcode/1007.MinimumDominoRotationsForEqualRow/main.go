package main

// Complexity:
// Time O(N) and Space O(1) where N is the length of tops and bottoms.
func minDominoRotations(tops []int, bottoms []int) int {
	result := minDominoRotationsOfTarget(tops, bottoms, tops[0])
	if result != -1 || tops[0] == bottoms[0] {
		return result
	}
	return minDominoRotationsOfTarget(tops, bottoms, bottoms[0])
}

func minDominoRotationsOfTarget(tops []int, bottoms []int, target int) int {
	rotateTop := 0
	rotateBottom := 0
	for i := range tops {
		if tops[i] != target && bottoms[i] != target {
			return -1
		}
		if tops[i] != target {
			rotateTop++
		}
		if bottoms[i] != target {
			rotateBottom++
		}
	}
	return min(rotateTop, rotateBottom)
}
