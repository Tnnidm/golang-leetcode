package majorityElement

// // HashMap方法，时间复杂度O(n)，空间复杂度O(n)
// func majorityElement(nums []int) int {
// 	stat := make(map[int]int)
// 	threshold := len(nums) / 2
// 	for _, num := range nums {
// 		if _, ok := stat[num]; ok {
// 			stat[num] += 1
// 		} else {
// 			stat[num] = 1
// 		}
// 		if stat[num] > threshold {
// 			return num
// 		}
// 	}
// 	return 0
// }

// // 排序法，原理是排序数组中下标为 ⌊n/2⌋ 的元素一定为题中定义的众数，时间复杂度是O(nlogn),空间复杂度是O(logn)
// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)/2]
// }

// Boyer-Moore 投票算法，本质上就是抵消，时间复杂度O(n)，空间复杂度O(1)
func majorityElement(nums []int) int {
	var candidate int
	var count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}
