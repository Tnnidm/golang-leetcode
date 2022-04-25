package MagicDictionary

type MagicDictionary struct {
	isEnd    bool
	children [26](*MagicDictionary)
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (head *MagicDictionary) Insert(word string) {
	temp := head
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			temp.children[index] = &MagicDictionary{}
			temp = temp.children[index]
		}
	}
	temp.isEnd = true
}

func (head *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		head.Insert(word)
	}
}

func (head *MagicDictionary) Search(searchWord string) bool {
	if len(searchWord) == 0 {
		return head.isEnd
	}
	index := searchWord[0] - 'a'
	if head.children[index] != nil && head.children[index].Search(searchWord[1:]) {
		return true
	}
	for i := 0; i < 26; i++ {
		if i != int(index) && head.children[i] != nil && head.children[i].secondarySearch(searchWord[1:]) {
			return true
		}
	}
	return false
}

func (head *MagicDictionary) secondarySearch(searchWord string) bool {
	temp := head
	for i := 0; i < len(searchWord); i++ {
		index := searchWord[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}
	return temp.isEnd
}
