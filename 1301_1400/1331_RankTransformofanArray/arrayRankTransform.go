package arrayRankTransform

import "sort"

func arrayRankTransform(arr []int) []int {
	arrLen := len(arr)

	result := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		result[i] = arr[i]
	}

	sort.Ints(arr)
	rankMap := make(map[int]int, arrLen)
	index := 1
	for i := 0; i < arrLen; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			rankMap[arr[i]] = index
			index++
		}
	}

	for i := 0; i < arrLen; i++ {
		result[i] = rankMap[result[i]]
	}
	return result
}
