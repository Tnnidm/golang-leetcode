package combinationSum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)
	sort.Ints(candidates)
	var backtracking func(start, tarVal int)
	backtracking = func(start, tarVal int) {
		if tarVal == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > tarVal {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtracking(i+1, tarVal-candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtracking(0, target)
	return res
}
