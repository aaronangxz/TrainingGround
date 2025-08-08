package main

func soupServings(n int) float64 {
	// probability is monotonicaly increasing towards 1
	if n > 5000 { // for n=5000 prob = 0.9999967386599964 which is within 10^-5 precision == 1
		return 1.0
	}
	return soupServings_recursion(n, n)
}

var soupServings_cache = make(map[[2]int]float64)

func soupServings_recursion(a, b int) float64 {
	res := 0.0
	// caching
	if res, exists := soupServings_cache[[2]int{a, b}]; exists {
		return res
	}

	cases := []struct{ froma, fromb int }{{100, 0}, {75, 25}, {50, 50}, {25, 75}}
	for _, c := range cases {
		newa := max(a-c.froma, 0)
		newb := max(b-c.fromb, 0)

		prob := 0.0
		if newa == 0 { // a empty
			if newb > 0 { // a before b
				prob = 1
			} else {
				prob = 0.5 // a and b simultaneously == 0
			}
		} else if newb == 0 { // newb is 0, newa>0
			prob = 0
		} else if newb != 0 {
			prob = soupServings_recursion(newa, newb) //we continue recursion
		}
		// cache our result
		soupServings_cache[[2]int{newa, newb}] = prob
		//
		res += prob / 4
	}
	return res
}
