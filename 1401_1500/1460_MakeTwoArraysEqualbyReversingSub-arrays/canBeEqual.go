package canBeEqual

import "sort"

func canBeEqual(target []int, arr []int) bool {
	targetLen := len(target)
	arrLen := len(arr)
	if targetLen != arrLen {
		return false
	}
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < targetLen; i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
