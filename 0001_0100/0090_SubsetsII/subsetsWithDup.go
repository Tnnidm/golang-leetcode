package subsetsWithDup

import "sort"

// func subsetsWithDup(nums []int) [][]int {
// 	numsAre := make([]int, 0)
// 	numsCount := make([]int, 21)
// 	for _, num := range nums {
// 		if numsCount[num+10] == 0 {
// 			numsAre = append(numsAre, num)
// 		}
// 		numsCount[num+10]++
// 	}
// 	numsNum := len(numsAre)
// 	result := [][]int{}
// 	combination := []int{}
// 	var dfs func(index int)
// 	dfs = func(index int) {
// 		if index == numsNum {
// 			result = append(result, append([]int{}, combination...))
// 			return
// 		}
// 		dfs(index + 1)
// 		combinationLen := len(combination)
// 		for i := 0; i < numsCount[index]; i++ {
// 			combination = append(combination, numsAre[index])
// 		}
// 		combination = combination[:combinationLen]
// 	}
// 	dfs(0)
// 	return result
// }

// 一种复杂度更低的方法
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	combination := []int{}
	result := [][]int{}
	var dfs func(bool, int)
	numsLen := len(nums)
	dfs = func(choosePre bool, index int) {
		if index == numsLen {
			result = append(result, append([]int{}, combination...))
			return
		}
		dfs(false, index+1)
		if !choosePre && index > 0 && nums[index-1] == nums[index] {
			return
		}
		combination = append(combination, nums[index])
		dfs(true, index+1)
		combination = combination[:len(combination)-1]
	}
	dfs(false, 0)
	return result
}
