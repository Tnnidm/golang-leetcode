package randomizedSet

import "math/rand"

type RandomizedSet struct {
	hashMapIndex   map[int]int
	valueContainer []int
	tailIndex      int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}, -1}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rdset *RandomizedSet) Insert(val int) bool {
	if _, ok := rdset.hashMapIndex[val]; !ok {
		rdset.valueContainer = append(rdset.valueContainer, val)
		rdset.tailIndex++
		rdset.hashMapIndex[val] = rdset.tailIndex
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rdset *RandomizedSet) Remove(val int) bool {
	if index, ok := rdset.hashMapIndex[val]; ok {
		if index != rdset.tailIndex {
			rdset.valueContainer[index], rdset.valueContainer[rdset.tailIndex] = rdset.valueContainer[rdset.tailIndex], rdset.valueContainer[index]
			rdset.hashMapIndex[rdset.valueContainer[index]] = index
		}
		delete(rdset.hashMapIndex, rdset.valueContainer[rdset.tailIndex])
		rdset.valueContainer = rdset.valueContainer[:rdset.tailIndex]
		rdset.tailIndex--
		return true
	}
	return false
}

/** Get a random element from the set. */
func (rdset *RandomizedSet) GetRandom() int {
	return rdset.valueContainer[rand.Intn(rdset.tailIndex+1)]
}
