package replaceWords

import "strings"

type Trie struct {
	isEnd    bool
	children [26](*Trie)
}

func (head *Trie) Insert(word string) {
	temp := head
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			temp.children[index] = &Trie{}
			temp = temp.children[index]
		}
	}
	temp.isEnd = true
}

func (head *Trie) SearchPrefix(word string) string {
	temp := head
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
			if temp.isEnd {
				return word[:i+1]
			}
		} else {
			break
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	head := &Trie{}
	for _, word := range dictionary {
		head.Insert(word)
	}
	words := strings.Fields(sentence)
	for i := range words {
		words[i] = head.SearchPrefix(words[i])
	}
	return strings.Join(words, " ")
}
