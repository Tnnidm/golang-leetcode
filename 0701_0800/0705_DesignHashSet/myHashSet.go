package main

import "fmt"

type MyHashSet struct {
	modHash int
	hashMap [][]int
}

func Constructor() MyHashSet {
	hs := MyHashSet{}
	hs.modHash = 57
	hs.hashMap = make([][]int, hs.modHash)
	return hs
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.hashMap[key%this.modHash] = append(this.hashMap[key%this.modHash], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	for i := range this.hashMap[key%this.modHash] {
		if this.hashMap[key%this.modHash][i] == key {
			this.hashMap[key%this.modHash] = append(this.hashMap[key%this.modHash][:i], this.hashMap[key%this.modHash][i+1:]...)
			break
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for i := range this.hashMap[key%this.modHash] {
		if this.hashMap[key%this.modHash][i] == key {
			return true
		}
	}
	return false
}

func main() {
	myHashSet := Constructor()
	myHashSet.Add(1) // set = [1]
	fmt.Println(myHashSet.hashMap[:3])
	myHashSet.Add(2) // set = [1, 2]
	fmt.Println(myHashSet.hashMap[:3])
	myHashSet.Contains(1) // return True
	myHashSet.Contains(3) // return False, (not found)
	myHashSet.Add(2)      // set = [1, 2]
	fmt.Println(myHashSet.hashMap[:3])
	myHashSet.Contains(2) // return True
	myHashSet.Remove(2)   // set = [1]
	fmt.Println(myHashSet.hashMap[:3])
	myHashSet.Contains(2) // return False, (already removed)
}
