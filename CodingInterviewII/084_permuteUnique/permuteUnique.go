package permuteUnique

// // 方法1: 使用map去除重复
// func permuteUnique(nums []int) [][]int {
// 	numsLen := len(nums)
// 	result := [][]int{}
// 	var dfs func(index int)
// 	dfs = func(index int) {
// 		if index == numsLen-1 {
// 			result = append(result, append([]int{}, nums...))
// 			return
// 		}
// 		dfs(index + 1)
// 		exits := map[int]bool{}
// 		for i := index + 1; i < numsLen; i++ {
// 			if nums[i] != nums[index] && !exits[nums[i]] {
// 				nums[i], nums[index] = nums[index], nums[i]
// 				dfs(index + 1)
// 				nums[i], nums[index] = nums[index], nums[i]
// 				exits[nums[i]] = true
// 			}
// 		}
// 	}
// 	dfs(0)
// 	return result
// }

// 方法1: 鉴于这道题num大小范围有限，把map换成[]bool能省很多空间
func permuteUnique(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == numsLen-1 {
			result = append(result, append([]int{}, nums...))
			return
		}
		dfs(index + 1)
		var exits [21]bool
		for i := index + 1; i < numsLen; i++ {
			if nums[i] != nums[index] && !exits[nums[i]+10] {
				nums[i], nums[index] = nums[index], nums[i]
				dfs(index + 1)
				nums[i], nums[index] = nums[index], nums[i]
				exits[nums[i]+10] = true
			}
		}
	}
	dfs(0)
	return result
}
