package MapSum

type MapSum struct {
	key      int
	children [26](*MapSum)
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	temp := this
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] == nil {
			temp.children[index] = &MapSum{}
		}
		temp = temp.children[index]
	}
	temp.key = val
}

func (this *MapSum) Sum(prefix string) int {
	temp := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if temp.children[index] == nil {
			return 0
		}
		temp = temp.children[index]
	}
	result := 0
	temp.dfs(&result)
	return result
}

func (this *MapSum) dfs(sum *int) {
	*sum += this.key
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			this.children[i].dfs(sum)
		}
	}
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
