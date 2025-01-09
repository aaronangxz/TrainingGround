package main

type TrieNode struct {
	Children map[string]*TrieNode
	Count    int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[string]*TrieNode), Count: 0,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) add(str string) {
	curr := t.root
	for _, s := range str {
		sStr := string(s)
		// If the current char is not in the TrieNode children, add it
		if _, ok := curr.Children[sStr]; !ok {
			curr.Children[sStr] = NewTrieNode()
		}
		// Then / Else, move to the next TrieNode
		curr = curr.Children[sStr]
		// While also incrementing the total count of nodes
		curr.Count++
	}
	// first str = "attention"
	// Children = a (1) -> t (1) -> t (1) -> e (1) -> n (1) -> t (1) -> i (1) -> o (1) -> n (1)
	// second str = "attend"
	// Children = a (2) -> t (2) -> t (2) -> e (2) -> n (2) -> d (1)
	// Since "atten" is the common prefix, these have count of 2
	//			root
	//			/
	//			a (2)
	//			/
	//			t (2)
	//			/
	//			t (2)
	//			/
	//			e (2)
	//			/
	//			n (2)
	//			/	\
	//		t (1)	d (1)
	//		/
	//		i (1)
	//		/
	//		o (1)
	//		/
	//		n (1)
}

func (t *Trie) count(prefix string) int {
	curr := t.root
	for _, s := range prefix {
		sStr := string(s)
		// If we cannot even find the prefix in the Trie, that means no words have this prefix
		if _, ok := curr.Children[sStr]; !ok {
			return 0
		}
		// Otherwise move to the next TrieNode
		curr = curr.Children[sStr]
	}
	// Found the end of prefix, return the count of all words with this prefix
	// i.e. prefix = "at"
	//			root
	//			/
	//			a (2)
	//			/
	//			t (2)	 <- end of prefix, return 2
	//			/
	//			t (2)
	//			/
	//			e (2)
	//			/
	//			n (2)
	//			/	\
	//		t (1)	d (1)
	//		/
	//		i (1)
	//		/
	//		o (1)
	//		/
	//		n (1)
	return curr.Count
}

func prefixCount(words []string, pref string) int {
	trie := NewTrie()

	for _, w := range words {
		// Only add to trie if word is at least as long as prefix
		if len(w) >= len(pref) {
			trie.add(w)
		}
	}
	return trie.count(pref)
}
