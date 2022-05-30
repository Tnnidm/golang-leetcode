package subsets

func subsets(nums []int) [][]int {
	result := [][]int{}
	numsLen := len(nums)
	combination := []int{}
	var dfs func(int)
	dfs = func(index int) {
		if index == numsLen {
			result = append(result, append([]int{}, combination...))
			return
		}
		dfs(index + 1)
		combination = append(combination, nums[index])
		dfs(index + 1)
		combination = combination[:len(combination)-1]
	}
	dfs(0)
	return result
}
