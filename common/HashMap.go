package common

import "fmt"

const (
	defaultMapSize = 10
)

type HashNode struct {
	Key   int
	Value int
	Next  *HashNode
}

func NewHashNode(key, value int) *HashNode {
	return &HashNode{
		Key:   key,
		Value: value,
		Next:  nil,
	}
}

type HashMap struct {
	Table []*HashNode
	Size  int
}

func NewHashMap(n int) *HashMap {
	return &HashMap{Table: make([]*HashNode, n), Size: n}
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
	h.Size++
}

func (h *HashMap) Get(key int) int {
	k := h.hashFunction(key)

	if h.Table[k] != nil {
		curr := h.Table[k]
		for curr.Next != nil {
			if curr.Key == key {
				return curr.Value
			}
			curr = curr.Next
		}
	}
	return -1
}

func (h *HashMap) Delete(key int) {
	k := h.hashFunction(key)

	if h.Table[k].Key == key {
		h.Size--
		h.Table[k] = h.Table[k].Next
		return
	}

	curr := h.Table[k]
	for curr != nil {
		if curr.Next != nil && curr.Next.Key == key {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	h.Size--
}

func (h *HashMap) PrintHashMap() {
	for i, m := range h.Table {
		fmt.Println("Bucket", i)
		if m == nil {
			fmt.Println("Empty")
		} else {
			for m != nil {
				fmt.Printf("Key: %v, Value: %v\n", m.Key, m.Value)
				m = m.Next
			}
		}
		fmt.Println("")
	}
}
