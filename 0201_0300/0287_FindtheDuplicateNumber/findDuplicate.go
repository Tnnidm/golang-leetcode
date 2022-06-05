package findDuplicate

// 自地址法
func findDuplicate(nums []int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		nums[i] = 100000 * nums[i]
	}
	for i := 0; i < numsLen; i++ {
		if nums[nums[i]/100000]%100000 == 1 {
			return nums[i] / 100000
		}
		nums[nums[i]/100000]++
	}
	return 0
}

// // 快慢指针法
// // 难点是把他当作带环链表的思路
// func findDuplicate(nums []int) int {
// 	slow, fast := 0, 0
// 	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
// 	}
// 	slow = 0
// 	for slow != fast {
// 		slow = nums[slow]
// 		fast = nums[fast]
// 	}
// 	return slow
// }
