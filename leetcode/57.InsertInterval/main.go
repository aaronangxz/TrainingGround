package main

import "github.com/aaronangxz/TrainingGround/common"

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	for i, intv := range intervals {
		// end of newInterval < start of first interval
		// It means none of those in the back can be smaller than newInterval
		if newInterval[1] < intv[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > intv[1] {
			// start of newInterval > end of first interval
			// not reaching the newInterval yet. Just append
			res = append(res, intv)
		} else {
			// Time to merge
			// merged = [min of (intv start & newInterval start),max of (intv end & newInterval end)]
			newInterval[0] = common.Min(intv[0], newInterval[0])
			newInterval[1] = common.Max(intv[1], newInterval[1])
		}
	}
	// Finally, append the (modified) newInterval
	res = append(res, newInterval)
	return res
}
