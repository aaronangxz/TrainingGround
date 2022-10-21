package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int)
	var uniq []string

	//Iterate through words
	for _, w := range words {
		//if it's a new word (not seen in map)
		if _, ok := freq[w]; !ok {
			//save into unique string slice
			uniq = append(uniq, w)
		}
		//increment number of appearance
		freq[w]++
	}

	//sort the unique string slice with conditions
	sort.SliceStable(uniq, func(i, j int) bool {
		//if both words have same number of appearance, sort by lexicographical order
		if freq[uniq[i]] == freq[uniq[j]] {
			return uniq[i] < uniq[j]
		}
		//else sort by number of appearance
		return freq[uniq[i]] > freq[uniq[j]]
	})

	//only return the first kth words
	return uniq[:k]
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
