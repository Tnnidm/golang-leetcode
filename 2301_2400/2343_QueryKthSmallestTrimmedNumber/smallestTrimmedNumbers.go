package smallestTrimmedNumbers

import (
	"sort"
)

type numWithIndex struct {
	index int
	num   string
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	numsNum := len(nums)
	numsLen := len(nums[0])
	numWithIndexs := make([]numWithIndex, numsNum)
	for i := range nums {
		numWithIndexs[i].index = i
		numWithIndexs[i].num = nums[i]
	}
	result := make([]int, len(queries))
	for r := range result {
		sort.Slice(numWithIndexs, func(i, j int) bool {
			return numWithIndexs[i].num[numsLen-queries[r][1]:] < numWithIndexs[j].num[numsLen-queries[r][1]:] ||
				(numWithIndexs[i].num[numsLen-queries[r][1]:] == numWithIndexs[j].num[numsLen-queries[r][1]:] && numWithIndexs[i].index < numWithIndexs[j].index)
		})
		result[r] = numWithIndexs[queries[r][0]-1].index
	}
	return result
}
