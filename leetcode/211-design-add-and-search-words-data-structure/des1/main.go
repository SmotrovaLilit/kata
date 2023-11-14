package des1

type WordDictionary struct {
	child [26]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (t *WordDictionary) AddWord(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.child[idx] == nil {
			node.child[idx] = &WordDictionary{}
		}
		node = node.child[idx]
	}
	node.isEnd = true
}

func (t *WordDictionary) Search(word string) bool {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if word[i] != '.' {
			if node.child[idx] == nil {
				return false
			}
			node = node.child[idx]
			continue
		}
		for j := 0; j < len(node.child); j++ {
			if node.child[j] == nil {
				continue
			}
			if node.child[j].Search(word[i+1:]) {
				return true
			}
		}
		return false
	}
	return node.isEnd
}
