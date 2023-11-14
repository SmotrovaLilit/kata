package des1

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

// Insert inserts a word into the trie.
// Time complexity: O(m), where m is the length of the word.
// Space complexity: O(m).
func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		if node.child[word[i]-'a'] == nil {
			node.child[word[i]-'a'] = &Trie{}
		}
		node = node.child[word[i]-'a']
	}
	node.isEnd = true
}

// Search returns true if the word is in the trie.
// Time complexity: O(m), where m is the length of the word.
// Space complexity: O(1).
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix.
// Time complexity: O(m), where m is the length of the prefix.
// Space complexity: O(1).
func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}

func (t *Trie) searchPrefix(word string) *Trie {
	node := t
	for i := 0; i < len(word); i++ {
		if node.child[word[i]-'a'] != nil {
			node = node.child[word[i]-'a']
		} else {
			return nil
		}
	}
	return node
}
