package main

import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {

	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}

	for i := 0; i < len(s1); i++ {
		a := int(s1[i] - 'a')
		b := int(s2[i] - 'a')
		ra := find(a)
		rb := find(b)

		if ra != rb {
			if ra < rb {
				parent[rb] = ra
			} else {
				parent[ra] = rb
			}
		}
	}

	var sb strings.Builder
	for i := 0; i < len(baseStr); i++ {
		r := find(int(baseStr[i] - 'a'))
		sb.WriteByte(byte(r + 'a'))
	}
	return sb.String()
}
