package Trie

type Trie struct {
	children [26](*Trie)
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if this.children[word[i]-'a'] == nil {
			this.children[word[i]-'a'] = &Trie{}
		}
		this = this.children[word[i]-'a']
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if this.children[word[i]-'a'] == nil {
			return false
		}
		this = this.children[word[i]-'a']
	}
	return this.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if this.children[prefix[i]-'a'] == nil {
			return false
		}
		this = this.children[prefix[i]-'a']
	}
	return true
}
