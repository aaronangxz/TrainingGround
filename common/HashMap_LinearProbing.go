package common

import "fmt"

type HashMapLinearProbing struct {
	Table    []*HashNodeSingle
	Size     int //current size
	Capacity int //total size
}

func NewHashMapLinearProbing(n int) *HashMapLinearProbing {
	return &HashMapLinearProbing{Table: make([]*HashNodeSingle, n), Size: 0, Capacity: n}
}

func NewHashMapLinearProbingFromSlice(s map[int]int) *HashMapLinearProbing {
	m := NewHashMapLinearProbing(defaultMapSize)

	for key, val := range s {
		m.Insert(key, val)
	}

	return m
}

func (h *HashMapLinearProbing) hashFunction(key int) int {
	return key % h.Capacity
}

func (h *HashMapLinearProbing) Insert(key, value int) {
	k := h.hashFunction(key)

	//Increment key to find next empty slot
	for h.Table[k] != nil && h.Table[k].Key != key && h.Table[k].Key != -1 {
		k++
		k %= h.Capacity
	}

	if h.Table[k] == nil || h.Table[k].Key == -1 {
		h.Size++
		h.Table[k] = NewHashNodeSingle(key, value)
	}
}

func (h *HashMapLinearProbing) Get(key int) int {
	k := h.hashFunction(key)
	count := 0

	for h.Table[k] != nil {
		if count > h.Capacity {
			return -1
		}

		if h.Table[k].Key == key {
			fmt.Printf("Value of %v is %v\n", key, h.Table[k].Value)
			return h.Table[k].Value
		}
		k++
		k %= h.Capacity
	}
	return -1
}

func (h *HashMapLinearProbing) Delete(key int) {
	k := h.hashFunction(key)

	for h.Table[k] != nil {
		if h.Table[k].Key == key {
			h.Table[k] = NewHashNodeSingle(-1, -1)
			h.Size--
		}
		k++
		k %= h.Capacity
	}
}

func (h *HashMapLinearProbing) PrintHashMap() {
	for i, n := range h.Table {
		fmt.Printf("Bucket%v ", i)
		n.PrintNode()
		fmt.Println()
	}
}
