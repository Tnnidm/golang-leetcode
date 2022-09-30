package sortPeople

import "sort"

func sortPeople(names []string, heights []int) []string {
	heightToName := map[int]string{}
	for i := range names {
		heightToName[heights[i]] = names[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	for i := range heights {
		names[i] = heightToName[heights[i]]
	}
	return names
}
