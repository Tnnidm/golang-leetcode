package merge

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	intervalsNum := len(intervals)
	result := [][]int{intervals[0]}
	lastIndex := 0
	for i := 1; i < intervalsNum; i++ {
		if result[lastIndex][1] >= intervals[i][0] {
			result[lastIndex][1] = max(result[lastIndex][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
			lastIndex++
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
