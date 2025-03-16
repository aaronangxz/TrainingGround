package main

type Node struct {
	Prev, Next *Node
	Key, Value int
}

type LRUCache struct {
	Head, Tail *Node
	Mapping    map[int]*Node
	Capacity   int
}

func Constructor(capacity int) LRUCache {
	h := &Node{
		Key:   0,
		Value: 0,
	}
	t := &Node{
		Key:   0,
		Value: 0,
	}
	h.Next = t
	t.Prev = h
	return LRUCache{
		Head:     h,
		Tail:     t,
		Mapping:  make(map[int]*Node),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, found := this.Mapping[key]; found {
		delete(this.Mapping, key)
		// Head -- ... -- Some Node A -x- Curr Node -x- Some Node B -- ...
		// Head -- ... -- Some Node A -- Some Node B -- ...
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev

		this.Put(key, n.Value)
		return n.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// If found, delete from map and node
	if n, found := this.Mapping[key]; found {
		delete(this.Mapping, key)
		// Head -- ... -- Some Node A -x- Curr Node -x- Some Node B -- ...
		// Head -- ... -- Some Node A -- Some Node B -- ...
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}

	// If going to exceed capacity in next Put, evict LRU
	if len(this.Mapping) == this.Capacity {
		LRU := this.Tail.Prev
		delete(this.Mapping, LRU.Key)
		// Head -- ... -- 2nd LRU Node -x- Curr LRU Node -x- Tail -- Nil
		// Head -- ... -- New LRU Node -- Tail -- Nil
		LRU.Prev.Next = LRU.Next
		LRU.Next.Prev = LRU.Prev
	}

	// Add back to map with new Node pointer
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	this.Mapping[key] = newNode
	// Connect new Node to link with Head, before current MRU Node
	// Head -^- current MRU Node -- ... -- Tail
	// Head -- NewNode -- previous MRU Node -- ... -- Tail
	next := this.Head.Next
	this.Head.Next = newNode
	newNode.Prev = this.Head
	next.Prev = newNode
	newNode.Next = next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
