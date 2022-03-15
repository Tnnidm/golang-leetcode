package minNumber

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	numsNum := len(nums)
	numsStr := make([]string, numsNum)
	for i := 0; i < numsNum; i++ {
		numsStr[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(numsStr, func(i, j int) bool {
		temp1 := numsStr[i] + numsStr[j]
		temp2 := numsStr[j] + numsStr[i]
		tempLen := len(temp1)
		for k := 0; k < tempLen; k++ {
			if temp1[k] < temp2[k] {
				return true
			} else if temp1[k] > temp2[k] {
				return false
			}
		}
		return false
	})
	result := strings.Builder{}
	for i := 0; i < numsNum; i++ {
		result.WriteString(numsStr[i])
	}
	return result.String()
}
