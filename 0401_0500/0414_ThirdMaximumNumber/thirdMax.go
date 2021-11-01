package thirdMax

import (
	"math"
)

// // 思路就是排序以后找第三大的
// func thirdMax(nums []int) int {
// 	sort.Ints(nums)
// 	thirdMAx := nums[len(nums)-1]
// 	counter := 0
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i] != nums[i-1] {
// 			counter++
// 		}
// 		if counter == 2 {
// 			thirdMAx = nums[i-1]
// 			break
// 		}
// 	}
// 	return thirdMAx
// }

// 进阶：你能设计一个时间复杂度 O(n) 的解决方案吗？
func thirdMax(nums []int) int {
	firstMax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > thirdMax && nums[i] != firstMax && nums[i] != secondMax {
			thirdMax = nums[i]
		}
		if thirdMax > secondMax {
			thirdMax, secondMax = secondMax, thirdMax
		}
		if secondMax > firstMax {
			secondMax, firstMax = firstMax, secondMax
		}
	}
	if thirdMax != math.MinInt64 {
		return thirdMax
	} else {
		return firstMax
	}
}
