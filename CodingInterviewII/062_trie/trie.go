package trie

// // 用map来实现多叉树在时间上不够高效，因为前缀树都是小写字母，可以直接用数组简化成26叉树
// type Trie struct {
// 	childs map[byte](*Trie)
// 	isEnd  bool
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	return Trie{map[byte](*Trie){}, false}
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	for i := 0; i < len(word); i++ {
// 		if next, ok := this.childs[word[i]]; ok {
// 			this = next
// 		} else {
// 			this.childs[word[i]] = &Trie{map[byte](*Trie){}, false}
// 			this = this.childs[word[i]]
// 		}
// 	}
// 	this.isEnd = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	for i := 0; i < len(word); i++ {
// 		if next, ok := this.childs[word[i]]; ok {
// 			this = next
// 		} else {
// 			return false
// 		}
// 	}
// 	return this.isEnd
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	for i := 0; i < len(prefix); i++ {
// 		if next, ok := this.childs[prefix[i]]; ok {
// 			this = next
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// 用数组构建26叉树的做法
type Trie struct {
	end    bool
	childs [26](*Trie)
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if this.childs[index] != nil {
			this = this.childs[index]
		} else {
			this.childs[index] = &Trie{}
			this = this.childs[index]
		}
	}
	this.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if this.childs[index] != nil {
			this = this.childs[index]
		} else {
			return false
		}
	}
	return this.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if this.childs[index] != nil {
			this = this.childs[index]
		} else {
			return false
		}
	}
	return true
}
