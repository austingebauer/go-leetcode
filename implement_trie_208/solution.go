package implement_trie_208

type Trie struct {
	root *trieNode
}

type trieNode struct {
	char     byte
	end      bool
	children [26]*trieNode
}

// Returns a Trie
func Constructor() Trie {
	return Trie{
		root: &trieNode{},
	}
}

// Inserts a word into the trie
func (this *Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		slot := word[i] - 'a'

		// If the child doesn't exist, create it
		if curr.children[slot] == nil {
			curr.children[slot] = &trieNode{
				char: word[i],
			}
		}

		// Advance curr to the slot
		curr = curr.children[slot]
	}

	// Mark the last node as the end of an inserted word
	curr.end = true
}

// Returns true if the word is in the trie
func (this *Trie) Search(word string) bool {
	return this.search(word, true)
}

// Returns true if there is any word in the trie that starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
	return this.search(prefix, false)
}

func (this *Trie) search(word string, needsEnd bool) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		slot := word[i] - 'a'
		if curr.children[slot] == nil {
			return false
		}

		curr = curr.children[slot]
	}

	if needsEnd {
		return curr.end
	}

	return true
}
