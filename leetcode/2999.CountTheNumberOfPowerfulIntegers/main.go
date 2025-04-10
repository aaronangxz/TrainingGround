package main

import (
	"math"
	"strconv"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	lowerBound := strconv.FormatInt(start-1, 10)
	upperBound := strconv.FormatInt(finish, 10)

	return countValidNumbers(upperBound, s, limit) - countValidNumbers(lowerBound, s, limit)
}

func countValidNumbers(x string, suffix string, maxDigit int) int64 {
	lengthX := len(x)
	lengthSuffix := len(suffix)

	if lengthX < lengthSuffix {
		return 0
	}

	if lengthX == lengthSuffix {
		if x >= suffix {
			return 1
		}
		return 0
	}

	prefixLen := lengthX - lengthSuffix
	var total int64 = 0
	suffixInX := x[prefixLen:]

	for i := 0; i < prefixLen; i++ {
		currentDigit := int(x[i] - '0')

		if currentDigit > maxDigit {
			total += int64(math.Pow(float64(maxDigit+1), float64(prefixLen-i)))
			return total
		}

		total += int64(currentDigit) * int64(math.Pow(float64(maxDigit+1), float64(prefixLen-i-1)))
	}

	if suffixInX >= suffix {
		total++
	}

	return total
}
