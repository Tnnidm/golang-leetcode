package subsets

func subsets(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	temp := []int{}
	var dfs func(int)
	dfs = func(index int) {
		if index == numsLen {
			result = append(result, append([]int{}, temp...))
			return
		}
		dfs(index + 1)
		temp = append(temp, nums[index])
		dfs(index + 1)
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return result
}
