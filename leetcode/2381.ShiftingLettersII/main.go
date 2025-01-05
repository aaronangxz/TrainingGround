package main

func shiftingLetters(s string, shifts [][]int) string {
	// prefix sum slice to keep track of diff
	shiftMap := make([]int, len(s)+1)

	for _, shift := range shifts {
		start := shift[0]
		end := shift[1]
		direction := shift[2]

		// If it is increasing
		// start index of the range +1
		// next of end index of the range -1
		// creating a 'boundary' for the affected range
		// when we iterate through, it will 'offset' away the +1 from the previous range
		if direction == 1 {
			shiftMap[start]++
			shiftMap[end+1]--
		} else {
			shiftMap[start]--
			shiftMap[end+1]++
		}
	}

	diff := 0
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		// diff is the rolling sum from the previous index
		// we can do that because we did the 'offset' earlier^
		diff += shiftMap[i]
		// modulo to achieve wrap-around a-z
		shift := (int(s[i]-'a') + diff) % 26
		// it can be negative (decrement direction), offset back to positive
		if shift < 0 {
			shift += 26
		}
		// The total shifts from 'a'
		res[i] = byte(shift) + 'a'
	}
	return string(res)
}
