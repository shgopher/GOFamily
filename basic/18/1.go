package main

type Trie struct {
	isWord   bool
	children [26]*Trie
}

// 返回一个新的
func NewTrie() *Trie {
	return &Trie{}
}

// 插入数据
func (t *Trie) Insert(word string) {
	cur := t
	for i, c := range word {
		n := c - 'a'

		if cur.children[n] == nil {
			cur.children[n] = &Trie{}

		}
		cur = cur.children[n]
		if i == len(word)-1 {
			cur.isWord = true
		}

	}
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return cur.isWord
}

// Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) (bool, []string) {
	cur := t
	var result []string
	for _, c := range prefix {
		n := c - 'a'
		if cur.children[n] == nil {
			return false, nil
		}
		cur = cur.children[n]
		result = append(result, cur)
	}
	return true, result
}
