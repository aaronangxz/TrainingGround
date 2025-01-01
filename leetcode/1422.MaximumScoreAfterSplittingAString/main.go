package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"math"
)

func maxScore(s string) int {
	zeros, ones, best := 0, 0, -math.MaxInt32
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == "0" {
			zeros++
		} else {
			ones++
		}
		best = common.Max(best, zeros-ones)
	}

	if string(s[len(s)-1]) == "1" {
		ones++
	}
	return best + ones
}
