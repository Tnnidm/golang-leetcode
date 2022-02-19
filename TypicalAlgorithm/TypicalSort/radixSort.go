package typicalSort

import "sort"

func radixSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	mod := 10
	bucket := make([][]int, mod)
	for i := 0; i < mod; i++ {
		bucket[i] = []int{}
	}
	for i := 0; i < arrLen; i++ {
		bucket[arr[i]%mod] = append(bucket[arr[i]%mod], arr[i])
	}
	for i := 0; i < mod; i++ {
		sort.Ints(bucket[i])
	}
	index := 0
	for i := 0; i < mod; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			arr[index] = bucket[i][j]
			index++
		}
	}
	return arr
}
