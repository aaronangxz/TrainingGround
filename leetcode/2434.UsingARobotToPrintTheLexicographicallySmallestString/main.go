package main

func robotWithString(s string) string {
	a := 'a'
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-a]++
	}

	var stack []rune
	var result []rune

	hasSmaller := func(top int) bool {
		for i := 0; i < top; i++ {
			if freq[i] > 0 {
				return true
			}
		}
		return false
	}

	for _, ch := range s {
		idx := ch - a
		freq[idx]--
		stack = append(stack, ch)

		for len(stack) > 0 && !hasSmaller(int(stack[len(stack)-1]-a)) {
			result = append(result, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return string(result)
}
