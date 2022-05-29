package permute

func permute(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	var dfs func(int)
	dfs = func(index int) {
		if index == numsLen-1 {
			result = append(result, append([]int{}, nums...))
			return
		}
		dfs(index + 1)
		for i := index + 1; i < numsLen; i++ {
			nums[index], nums[i] = nums[i], nums[index]
			dfs(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	dfs(0)
	return result
}
