package merge

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := [][]int{intervals[0]}
	index := 0
	for i := 1; i < len(intervals); i++ {
		if result[index][1] >= intervals[i][0] {
			result[index][1] = max(result[index][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
			index++
		}
	}
	return result
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
