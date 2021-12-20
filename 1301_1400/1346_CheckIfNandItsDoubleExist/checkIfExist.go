package checkIfExist

import "sort"

func checkIfExist(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})
	arrLen := len(arr)
	doubleMap := make(map[int]struct{}, arrLen)
	for i := 0; i < arrLen; i++ {
		doubleMap[2*arr[i]] = struct{}{}
		if _, ok := doubleMap[arr[i]]; ok {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
