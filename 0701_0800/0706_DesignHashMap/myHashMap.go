package MyHashMap

type MyHashMap struct {
	mod     int
	hashMap [][][]int
}

func Constructor() MyHashMap {
	hs := MyHashMap{769, make([][][]int, 769)}
	return hs
}

func (this *MyHashMap) Put(key int, value int) {
	hashValue := key % this.mod
	notExitFlag := true
	for i := range this.hashMap[hashValue] {
		if this.hashMap[hashValue][i][0] == key {
			this.hashMap[hashValue][i][1] = value
			notExitFlag = false
		}
	}
	if notExitFlag {
		this.hashMap[hashValue] = append(this.hashMap[hashValue], []int{key, value})
	}
}

func (this *MyHashMap) Get(key int) int {
	hashValue := key % this.mod
	for i := range this.hashMap[hashValue] {
		if this.hashMap[hashValue][i][0] == key {
			return this.hashMap[hashValue][i][1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hashValue := key % this.mod
	for i := range this.hashMap[hashValue] {
		if this.hashMap[hashValue][i][0] == key {
			this.hashMap[hashValue] = append(this.hashMap[hashValue][:i], this.hashMap[hashValue][i+1:]...)
			break
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
