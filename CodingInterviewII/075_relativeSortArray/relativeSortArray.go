package relativeSortArray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Len := len(arr2)
	countMap := make(map[int]int, arr2Len)
	for i := 0; i < arr2Len; i++ {
		countMap[arr2[i]] = i - arr2Len
	}
	sort.Slice(arr1, func(i, j int) bool {
		x := arr1[i]
		if rank, ok := countMap[x]; ok {
			x = rank
		}
		y := arr1[j]
		if rank, ok := countMap[y]; ok {
			y = rank
		}
		return x <= y
	})
	return arr1
}
