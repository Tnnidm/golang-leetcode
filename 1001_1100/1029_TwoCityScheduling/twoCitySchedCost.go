package twoCitySchedCost

import "sort"

func twoCitySchedCost(costs [][]int) int {
	costsLen := len(costs)
	sum := 0
	if costsLen == 0 {
		return sum
	}
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] < costs[j][1]-costs[j][0]
	})

	for i := 0; i < costsLen/2; i++ {
		sum += costs[i][1]
	}
	for i := costsLen / 2; i < costsLen; i++ {
		sum += costs[i][0]
	}
	return sum
}
