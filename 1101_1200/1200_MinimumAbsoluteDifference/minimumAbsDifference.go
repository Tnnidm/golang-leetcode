package minimumAbsDifference

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := 2000000
	minDiffSet := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minDiff {
			minDiff = arr[i+1] - arr[i]
			minDiffSet = [][]int{}
			minDiffSet = append(minDiffSet, []int{arr[i], arr[i+1]})
		} else if arr[i+1]-arr[i] == minDiff {
			minDiffSet = append(minDiffSet, []int{arr[i], arr[i+1]})
		}
	}
	return minDiffSet
}
