package findRelativeRanks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	scoreLen := len(score)
	value2IndexMap := make(map[int]int, scoreLen)
	for i := 0; i < scoreLen; i++ {
		value2IndexMap[score[i]] = i
	}
	sort.Ints(score)
	result := make([]string, scoreLen)
	for i := 0; i < scoreLen; i++ {
		if i == scoreLen-3 {
			result[value2IndexMap[score[i]]] = "Bronze Medal"
		} else if i == scoreLen-2 {
			result[value2IndexMap[score[i]]] = "Silver Medal"
		} else if i == scoreLen-1 {
			result[value2IndexMap[score[i]]] = "Gold Medal"
		} else {
			result[value2IndexMap[score[i]]] = strconv.Itoa(scoreLen - i)
		}
	}
	return result
}
