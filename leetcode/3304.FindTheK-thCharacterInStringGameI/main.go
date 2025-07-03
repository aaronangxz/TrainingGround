package main

func kthCharacter(k int) byte {
	r := []rune{'a'}
	for len(r) < k {
		n := len(r)
		for i := 0; i < n; i++ {
			r = append(r, r[i]+1)
			if r[len(r)-1] > 122 {
				r[len(r)-1] = 97
			}
		}
	}
	return byte(r[k-1])
}
