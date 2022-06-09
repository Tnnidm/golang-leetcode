package reconstructQueue

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	peopleNum := len(people)
	result := make([][]int, peopleNum)
	for i := 0; i < peopleNum; i++ {
		result[i] = []int{}
	}
	decreaseNum := 0
	result[people[0][1]] = append(result[people[0][1]], people[0]...)
	for i := 1; i < peopleNum; i++ {
		if people[i][0] == people[i-1][0] {
			decreaseNum++
		} else {
			decreaseNum = 0
		}
		emptyIndex := -1
		for j := 0; j < peopleNum; j++ {
			if len(result[j]) == 0 {
				emptyIndex++
				if emptyIndex == people[i][1]-decreaseNum {
					result[j] = append(result[j], people[i]...)
					break
				}
			}
		}
	}
	return result
}
