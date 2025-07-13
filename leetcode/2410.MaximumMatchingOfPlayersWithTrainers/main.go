package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	// sort
	sort.Ints(players)
	sort.Ints(trainers)

	tInd, res := 0, 0           // # trainer index
	for pInd := range players { // player index
		// while block; search for valid pair
		for tInd < len(trainers) && players[pInd] > trainers[tInd] {
			tInd++
		}
		// stop if reached end of trainers slice
		if tInd >= len(trainers) {
			break
		}

		// if match is ok
		if players[pInd] <= trainers[tInd] {
			tInd++
			res++
		}
	}
	return res
}
