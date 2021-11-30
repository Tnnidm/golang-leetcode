package isMonotonic

import "sort"

// func isMonotonic(nums []int) bool {
// 	inc, dec := true, true
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			inc = false
// 		}
// 		if nums[i] < nums[i+1] {
// 			dec = false
// 		}
// 	}
// 	return inc || dec
// }

// 还有一种直接调包的方法
func isMonotonic(nums []int) bool {
	return sort.IntsAreSorted(nums) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}
