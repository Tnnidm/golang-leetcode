package findMaximumXOR

type Trie struct {
	children [2](*Trie)
}

func (head *Trie) Insert(num int) {
	temp := head
	for i := 31; i >= 0; i-- {
		index := num >> i & 1
		if temp.children[index] == nil {
			temp.children[index] = &Trie{}
		}
		temp = temp.children[index]
	}
}

func (head *Trie) maxXorForANum(num int) int {
	temp := head
	result := 0
	for i := 31; i >= 0; i-- {
		index := (num >> i) & 1
		if temp.children[1-index] != nil {
			temp = temp.children[1-index]
			result += 1 << i
		} else {
			temp = temp.children[index]
		}
	}
	return result
}

func findMaximumXOR(nums []int) int {
	head := &Trie{}
	for _, num := range nums {
		head.Insert(num)
	}
	maxXor := 0
	for _, num := range nums {
		maxXor = max(maxXor, head.maxXorForANum(num))
	}
	return maxXor
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
