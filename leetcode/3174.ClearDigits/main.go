package main

func clearDigits(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			result = result[:len(result)-1]
		} else {
			result = result + string(ch)
		}
	}
	return result
}
