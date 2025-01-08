package main

func countPrefixSuffixPairs(words []string) int {
	ans := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				ans++
			}
		}
	}
	return ans
}

func isPrefixAndSuffix(str1, str2 string) bool {
	if len(str2) < len(str1) {
		return false
	}
	prefix, suffix := str2[:len(str1)], str2[len(str2)-len(str1):]
	return prefix == str1 && suffix == str1
}
