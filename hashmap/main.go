package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func main() {
	m := common.NewHashMap(10)
	insertion(m)
	get(m)
	batchCreate()
	batchCreateWithRandom()
}

func batchCreate() {
	maps := make(map[int]int)
	maps[1] = 1234
	maps[11] = 4321
	maps[33] = 3003
	maps[69] = 9999
	maps[999] = 10000
	m := common.NewHashMapFromSlice(maps)
	m.PrintHashMap()
}

func batchCreateWithRandom() {
	mm := common.MakeRandomIntKVMap(10)
	m := common.NewHashMapFromSlice(mm)
	m.PrintHashMap()
}

func insertion(m *common.HashMap) {
	m.Insert(5, 6969)
	m.Insert(6, 6969)
	m.Insert(26, 9999)
	m.Insert(36, 9999)
	m.Insert(25, 1234)
	m.PrintHashMap()
}

func get(m *common.HashMap) {
	fmt.Println(m.Get(26))
}
