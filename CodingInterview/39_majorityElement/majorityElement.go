package majorityElement

import "math/rand"

// // 方法1: 通过对某个元素计数来得到
// func majorityElement(nums []int) int {
// 	candidate := nums[0]
// 	count := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != candidate {
// 			count--
// 		} else {
// 			count++
// 		}
// 		if count == 0 {
// 			candidate = nums[i]
// 			count = 1
// 		}
// 	}
// 	return candidate
// }

// 方法2: 因为在这道题里面，中位数一定就是出现次数超过一半的数
// 所以可以使用类似快排的方法来求中位数
func majorityElement(nums []int) int {
	numsLen := len(nums)
	middle := numsLen / 2
	start := 0
	end := numsLen - 1
	index := partition(nums, start, end)
	for index != middle {
		if index > middle {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(nums, start, end)
	}
	return nums[index]
}

func partition(nums []int, start int, end int) int {
	if end == start {
		return start
	}
	index := rand.Intn(end-start) + start
	nums[index], nums[end] = nums[end], nums[index]
	small := start - 1
	for index = start; index < end; index++ {
		if nums[index] < nums[end] {
			small++
			if small != index {
				nums[index], nums[small] = nums[small], nums[index]
			}
		}
	}
	small++
	nums[small], nums[end] = nums[end], nums[small]
	return small
}
