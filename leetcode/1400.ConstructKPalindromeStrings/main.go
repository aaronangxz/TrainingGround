package main

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	charMap := make(map[string]int)

	for _, c := range s {
		charMap[string(c)]++
	}

	oddCount := 0
	for _, count := range charMap {
		if count%2 != 0 {
			oddCount++
		}
		if oddCount > k {
			return false
		}
	}
	return true
}
