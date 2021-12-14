package relativeSortArray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Len := len(arr2)
	arr2map := make(map[int]int, arr2Len)
	for i := 0; i < arr2Len; i++ {
		arr2map[arr2[i]] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		index1, ok1 := arr2map[arr1[i]]
		index2, ok2 := arr2map[arr1[j]]
		if ok1 {
			if ok2 {
				return index1 < index2
			} else {
				return true
			}
		} else {
			if ok2 {
				return false
			} else {
				return arr1[i] < arr1[j]
			}
		}
	})
	return arr1
}
