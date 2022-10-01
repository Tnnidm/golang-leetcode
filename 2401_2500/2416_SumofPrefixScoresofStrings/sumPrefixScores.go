package sumPrefixScores

type Trie struct {
	children [26](*Trie)
	count    int
}

func sumPrefixScores(words []string) []int {
	root := &Trie{}
	for _, word := range words {
		temp := root
		for i := range word {
			if temp.children[word[i]-'a'] == nil {
				temp.children[word[i]-'a'] = &Trie{}
			}
			temp = temp.children[word[i]-'a']
			temp.count++
		}
	}
	wordNum := len(words)
	result := make([]int, wordNum)
	for j, word := range words {
		temp := root
		for i := range word {
			temp = temp.children[word[i]-'a']
			result[j] += temp.count
		}
	}
	return result
}
