package main

type FindSumPairs struct {
	freq1 map[int]int
	nums2 []int
	freq2 map[int]int
}

// Complexity:
// Time O(N+M) and Space O(N+M) where N is the length of nums1
// and M is the length of nums2.
func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	freq1 := make(map[int]int)
	for _, num := range nums1 {
		freq1[num]++
	}

	freq2 := make(map[int]int)
	for _, num := range nums2 {
		freq2[num]++
	}

	return FindSumPairs{freq1, nums2, freq2}
}

// Complexity:
// Time O(1) and Space O(1).
func (this *FindSumPairs) Add(index int, val int) {
	old := this.nums2[index]
	new := old + val

	this.nums2[index] = new
	this.freq2[new]++
	this.freq2[old]--
	if this.freq2[old] == 0 {
		delete(this.freq2, old)
	}
}

// Complexity:
// Time O(min(N,M)) and Space O(1) where N is the length of nums1
// and M is the length of nums2.
func (this *FindSumPairs) Count(tot int) int {
	if len(this.freq1) < len(this.freq2) {
		return count(this.freq1, this.freq2, tot)
	} else {
		return count(this.freq2, this.freq1, tot)
	}
}

func count(freq1, freq2 map[int]int, tot int) int {
	result := 0
	for num1, count := range freq1 {
		num2 := tot - num1
		result += count * freq2[num2]
	}
	return result
}
