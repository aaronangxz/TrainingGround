package common

import (
	"log"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

const (
	defaultMapSize = 10
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func SwapAny(a interface{}, b interface{}) error {
	typeA := reflect.TypeOf(a)
	typeB := reflect.TypeOf(b)

	if typeA != typeB {
		log.Println("SwapAny | ", ErrorInvalidValue)
		return ErrorInvalidValue
	}

	valA := reflect.ValueOf(a).Elem()
	valB := reflect.ValueOf(b).Elem()
	tmp := valA.Interface()

	valA.Set(valB)
	valB.Set(reflect.ValueOf(tmp))

	return nil
}

func MakeRandomIntKVMap(size int) map[int]int {
	key := MakeSlice(size)
	val := MakeSlice(size * 1000)
	m := make(map[int]int)

	for i, k := range key {
		m[k] = val[i]
	}
	return m
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReverseSlice(s *[]int) {
	ReverseSliceRange(s, 0, len(*s)-1)
}

func ReverseSliceRange(s *[]int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
