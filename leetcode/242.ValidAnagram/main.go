package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m1[string(s[i])] += 1
		m1[string(t[i])] -= 1
	}
	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}
	return true
}
