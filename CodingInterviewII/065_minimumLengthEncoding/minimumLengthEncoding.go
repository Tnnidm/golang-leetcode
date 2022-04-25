package minimumLengthEncoding

type Trie struct {
	isLeaf   bool
	children [26](*Trie)
}

func (head *Trie) Insert(word string, result *int) {
	temp := head
	wordLen := len(word)
	for i := wordLen - 1; i >= 0; i-- {
		index := word[i] - 'a'
		if temp.children[index] == nil {
			temp.children[index] = &Trie{true, [26](*Trie){}}
			if temp.isLeaf {
				*result++
				temp.isLeaf = false
			} else {
				*result += wordLen - i + 1
			}
		}
		temp = temp.children[index]
	}
}

func (head *Trie) CountChildren() int {
	count := 0
	for i := 0; i < 26; i++ {
		if head.children[i] != nil {
			count++
		}
	}
	return count
}

func minimumLengthEncoding(words []string) int {
	head := &Trie{false, [26](*Trie){}}
	result := 0
	for _, word := range words {
		head.Insert(word, &result)
	}
	return result
}
