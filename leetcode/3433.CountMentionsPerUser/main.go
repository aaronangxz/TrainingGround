package main

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	byTime := make(map[int][][]string)
	for _, ev := range events {
		t, _ := strconv.Atoi(ev[1])
		byTime[t] = append(byTime[t], ev)
	}

	timestamps := make([]int, 0, len(byTime))
	for t := range byTime {
		timestamps = append(timestamps, t)
	}
	sort.Ints(timestamps)

	mentions := make([]int, numberOfUsers)
	isOnline := make([]bool, numberOfUsers)
	offlineUntil := make([]int, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		isOnline[i] = true
	}

	for _, t := range timestamps {
		evs := byTime[t]

		for i := 0; i < numberOfUsers; i++ {
			if !isOnline[i] && offlineUntil[i] <= t {
				isOnline[i] = true
				offlineUntil[i] = 0
			}
		}

		for _, ev := range evs {
			if ev[0] == "OFFLINE" {
				id, _ := strconv.Atoi(ev[2])
				isOnline[id] = false
				offlineUntil[id] = t + 60
			}
		}

		for _, ev := range evs {
			if ev[0] != "MESSAGE" {
				continue
			}
			tokens := strings.Fields(ev[2])
			for _, token := range tokens {
				if token == "ALL" {
					for i := 0; i < numberOfUsers; i++ {
						mentions[i]++
					}
				} else if token == "HERE" {
					for i := 0; i < numberOfUsers; i++ {
						if isOnline[i] {
							mentions[i]++
						}
					}
				} else if strings.HasPrefix(token, "id") {
					id, _ := strconv.Atoi(token[2:])
					if id >= 0 && id < numberOfUsers {
						mentions[id]++
					}
				}
			}
		}
	}

	return mentions
}
