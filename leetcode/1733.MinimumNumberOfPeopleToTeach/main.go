package main

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	blocks := (n + 63) / 64

	// masks: (m+1) x blocks, 1-based users
	masks := make([][]uint64, m+1)
	for i := 1; i <= m; i++ {
		masks[i] = make([]uint64, blocks)
		for _, lang := range languages[i-1] {
			idx := (lang - 1) / 64
			bit := uint((lang - 1) % 64)
			masks[i][idx] |= (1 << bit)
		}
	}

	// find candidates
	candidates := make(map[int]bool)
	for _, f := range friendships {
		u, v := f[0], f[1]
		canTalk := false
		for b := 0; b < blocks; b++ {
			if (masks[u][b] & masks[v][b]) != 0 {
				canTalk = true
				break
			}
		}
		if !canTalk {
			candidates[u] = true
			candidates[v] = true
		}
	}
	if len(candidates) == 0 {
		return 0
	}

	// count languages known among candidates by scanning their lists
	count := make([]int, n+1)
	for u := range candidates {
		for _, lang := range languages[u-1] {
			count[lang]++
		}
	}

	maxOverlap := 0
	for lang := 1; lang <= n; lang++ {
		if count[lang] > maxOverlap {
			maxOverlap = count[lang]
		}
	}

	return len(candidates) - maxOverlap
}
