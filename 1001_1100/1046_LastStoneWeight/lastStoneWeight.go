package lastStoneWeight

import "sort"

func lastStoneWeight(stones []int) int {
	stoneNum := len(stones)
	for stoneNum > 1 {
		sort.Ints(stones)
		if stones[stoneNum-2] == stones[stoneNum-1] {
			stones = stones[:stoneNum-2]
			stoneNum = stoneNum - 2
		} else {
			stones[stoneNum-2] = abs(stones[stoneNum-2] - stones[stoneNum-1])
			stones = stones[:stoneNum-1]
			stoneNum--
		}
	}
	if stoneNum == 1 {
		return stones[0]
	} else {
		return 0
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
