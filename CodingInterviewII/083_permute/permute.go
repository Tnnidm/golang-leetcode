package permute

func permute(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == numsLen-1 {
			result = append(result, append([]int{}, nums...))
			return
		}
		dfs(index + 1)
		for i := index + 1; i < numsLen; i++ {
			nums[i], nums[index] = nums[index], nums[i]
			dfs(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	dfs(0)
	return result
}
