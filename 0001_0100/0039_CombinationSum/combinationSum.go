package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	candidatesLen := len(candidates)
	combination := []int{}
	result := [][]int{}
	var dfs func(int, int)
	dfs = func(index int, target int) {
		if target == 0 {
			result = append(result, append(make([]int, 0, len(combination)), combination...))
			return
		}
		if index == candidatesLen {
			return
		}
		combinationLen := len(combination)
		for target >= 0 {
			dfs(index+1, target)
			target -= candidates[index]
			combination = append(combination, candidates[index])
		}
		combination = combination[:combinationLen]
	}
	dfs(0, target)
	return result
}
