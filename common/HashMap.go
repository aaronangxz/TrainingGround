package common

import "fmt"

type HashMap struct {
	Table []*HashNodeList
	Size  int
}

func NewHashMap(n int) *HashMap {
	return &HashMap{Table: make([]*HashNodeList, n), Size: n}
}

func NewHashMapFromSlice(s map[int]int) *HashMap {
	m := NewHashMap(defaultMapSize)

	for key, val := range s {
		m.Insert(key, val)
	}

	return m
}

func (h *HashMap) hashFunction(key int) int {
	return key % h.Size
}

func (h *HashMap) Insert(key, value int) {
	k := h.hashFunction(key)

	if h.Table[k] == nil {
		h.Table[k] = NewHashNode(key, value)
	} else {
		curr := h.Table[k]
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = NewHashNode(key, value)
	}
}

func (h *HashMap) Get(key int) int {
	k := h.hashFunction(key)

	if h.Table[k] != nil {
		curr := h.Table[k]
		for curr.Next != nil {
			if curr.Node.Key == key {
				fmt.Printf("Value of %v is %v\n", key, curr.Node.Value)
				return curr.Node.Value
			}
			curr = curr.Next
		}
	}
	return -1
}

func (h *HashMap) Delete(key int) {
	k := h.hashFunction(key)

	if h.Table[k].Node.Key == key {
		h.Table[k] = h.Table[k].Next
		return
	}

	curr := h.Table[k]
	for curr != nil {
		if curr.Next != nil && curr.Next.Node.Key == key {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
}

func (h *HashMap) PrintHashMap() {
	for i, m := range h.Table {
		fmt.Printf("Bucket %v :: ", i)
		m.PrintNodes()
	}
}
