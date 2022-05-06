package combinationSum

import "sort"

// // 方法1
// func combinationSum(candidates []int, target int) [][]int {
// 	candidatesNum := len(candidates)
// 	combination := []int{}
// 	result := [][]int{}
// 	var dfs func(int, int)
// 	dfs = func(index, target int) {
// 		if target == 0 {
// 			result = append(result, append([]int{}, combination...))
// 			return
// 		}
// 		if target < 0 || index == candidatesNum {
// 			return
// 		}
// 		combinationLen := len(combination)
// 		for target >= 0 {
// 			dfs(index+1, target)
// 			target -= candidates[index]
// 			combination = append(combination, candidates[index])
// 		}
// 		combination = combination[:combinationLen]
// 	}
// 	dfs(0, target)
// 	return result
// }

// 方法2
func combinationSum(candidates []int, target int) [][]int {
	combination, result := []int{}, [][]int{}
	sort.Ints(candidates)
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			result = append(result, append([]int(nil), combination...))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			combination = append(combination, candidates[i])
			dfs(i, target-candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	dfs(0, target)
	return result
}
