package main

import "math"

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	initial := (1 << n) - 1                                   // set first n bits to 1
	firstPlayer, secondPlayer = firstPlayer-1, secondPlayer-1 // convert to 0-based index

	var dfs func(state int, roundIdx int, lookingForMin bool, memo map[[2]int]int) int
	dfs = func(state int, roundIdx int, lookingForMin bool, memo map[[2]int]int) int {
		k := [2]int{state, roundIdx}
		if val, ok := memo[k]; ok {
			return val
		}

		nextStates := []int{0}
		for lo, hi := 0, 28; lo <= hi; lo, hi = lo+1, hi-1 {
			for state&(1<<lo) == 0 {
				lo++
			}
			for state&(1<<hi) == 0 {
				hi--
			}
			if lo > hi {
				break
			}

			if lo == firstPlayer && hi == secondPlayer {
				memo[k] = roundIdx
				return roundIdx
			} else if lo == firstPlayer || lo == secondPlayer || lo == hi { // in all those cases, lo should go to the next round
				setBit(nextStates, lo)
			} else if hi == firstPlayer || hi == secondPlayer { // in all those cases, hi should go to the next round
				setBit(nextStates, hi)
			} else { // we setBit both decisions to current states, either lo or hi win
				cpNextStates := make([]int, len(nextStates))
				copy(cpNextStates, nextStates)
				setBit(cpNextStates, hi)

				setBit(nextStates, lo)

				// merge the two sets of next states
				nextStates = append(nextStates, cpNextStates...)
			}
		}

		// run through all possible states
		res := math.MaxInt32
		if !lookingForMin {
			res = 0
		}
		for _, nextState := range nextStates {
			if lookingForMin {
				res = min(res, dfs(nextState, roundIdx+1, lookingForMin, memo))
			} else {
				res = max(res, dfs(nextState, roundIdx+1, lookingForMin, memo))
			}
		}

		memo[k] = res
		return res
	}

	return []int{
		dfs(initial, 1, true, make(map[[2]int]int)),
		dfs(initial, 1, false, make(map[[2]int]int)),
	}
}

func setBit(states []int, bitIdx int) {
	for i := range states {
		states[i] |= (1 << bitIdx)
	}
}
