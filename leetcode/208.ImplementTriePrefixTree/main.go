package main

type TrieNode struct {
	children []*TrieNode
	endHere  bool // Mark the end of a string branch
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		&TrieNode{
			children: make([]*TrieNode, 26),
			endHere:  false,
		},
	}
}

func (this *Trie) Insert(word string) {
	ptr := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		// If no node of this char exists, create a new one
		if ptr.children[ch-'a'] == nil {
			ptr.children[ch-'a'] = &TrieNode{
				children: make([]*TrieNode, 26),
				endHere:  false,
			}
		}
		// Update curr pointer to the latest node
		ptr = ptr.children[ch-'a']
	}
	// Mark as last node
	ptr.endHere = true
}

func (this *Trie) Search(word string) bool {
	ptr := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		// Such node not found
		if ptr.children[ch-'a'] == nil {
			return false
		}
		// Curr char found, continue to the next node
		ptr = ptr.children[ch-'a']
	}
	// Finish matching the word and arrived at a node
	// Only true if it ends here, otherwise it may just be a substring of a word branch
	// l -> e -> e -> t -> c -> o -> d -> e
	// Search("leet") ^, but actually its not the full word
	return ptr.endHere
}

func (this *Trie) StartsWith(prefix string) bool {
	// Same as Search, but ignore if the word string ends there
	ptr := this.root
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		if ptr.children[ch-'a'] == nil {
			return false
		}
		ptr = ptr.children[ch-'a']
	}
	return true
}
