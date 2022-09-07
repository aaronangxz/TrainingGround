package common

type HashMapQuadProbing struct {
	m HashMapLinearProbing
}

func NewHashMapQuadProbing(n int) *HashMapQuadProbing {
	return &HashMapQuadProbing{m: HashMapLinearProbing{
		Table: make([]*HashNodeSingle, n), Size: 0, Capacity: n}}
}

func NewHashMapQuadProbingFromSlice(s map[int]int) *HashMapQuadProbing {
	m := NewHashMapQuadProbing(defaultMapSize)

	for key, val := range s {
		m.Insert(key, val)
	}

	return m
}

func (h *HashMapQuadProbing) hashFunction(key int) int {
	return h.m.hashFunction(key)
}

func (h *HashMapQuadProbing) Insert(key, value int) {
	k := h.hashFunction(key)
	i := 0

	//Increment key to find next empty slot
	for h.m.Table[k] != nil && h.m.Table[k].Key != key && h.m.Table[k].Key != -1 && i < h.m.Capacity {
		i++
		k += i * i
		k %= h.m.Capacity
	}

	if h.m.Table[k] == nil || h.m.Table[k].Key == -1 {
		h.m.Size++
		h.m.Table[k] = NewHashNodeSingle(key, value)
	}
}

func (h *HashMapQuadProbing) Get(key int) int {
	k := h.hashFunction(key)
	count := 0
	i := 0

	for h.m.Table[k] != nil {
		if count > h.m.Capacity {
			return -1
		}

		if h.m.Table[k].Key == key {
			return h.m.Table[k].Value
		}
		i++
		k += i * i
		k %= h.m.Capacity
	}
	return -1
}

func (h *HashMapQuadProbing) Delete(key int) {
	k := h.hashFunction(key)
	i := 0

	for h.m.Table[k] != nil {
		if h.m.Table[k].Key == key {
			h.m.Table[k] = NewHashNodeSingle(-1, -1)
			h.m.Size--
			return
		}
		i++
		k += i * i
		k %= h.m.Capacity
	}
}

func (h *HashMapQuadProbing) PrintHashMap() {
	h.m.PrintHashMap()
}
