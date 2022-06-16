package common

import (
	"math/rand"
	"sort"
	"time"
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func MakeSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(size)
}

func MakeSortedSlice(size int) []int {
	s := MakeSlice(size)
	sort.Ints(s)
	return s
}
