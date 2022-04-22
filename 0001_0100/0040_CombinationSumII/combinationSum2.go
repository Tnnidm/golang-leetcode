package combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	candidatesLen := len(candidates)
	combination := []int{}
	result := [][]int{}
	var dfs func(int, int)
	dfs = func(index int, sum int) {
		if sum >= target {
			if sum == target {
				result = append(result, append([]int{}, combination...))
			}
			return
		}
		for i := index; i < candidatesLen; i++ {
			if i-1 >= index && candidates[i-1] == candidates[i] {
				continue
			}
			combination = append(combination, candidates[i])
			dfs(i+1, sum+candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	dfs(0, 0)
	return result
}
