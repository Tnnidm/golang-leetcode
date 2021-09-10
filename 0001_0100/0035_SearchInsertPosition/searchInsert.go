package searchInsert

// My method
// 执行用时：4 ms, 在所有 Go 提交中击败了81.14%的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了99.97%的用户
func searchInsert(nums []int, target int) int {
	index := 0
	numsLen := 0
	subLen := 0
	numsLen = len(nums)
	for {
		if numsLen == 1 {
			if target <= nums[0] {
				index += 0
			} else {
				index += 1
			}
			break
		} else {
			subLen = numsLen / 2
			subnums1 := nums[:subLen]
			subnums2 := nums[subLen:]
			if target <= subnums1[subLen-1] {
				nums = subnums1
				numsLen = subLen
			} else if target > subnums2[0] {
				nums = subnums2
				numsLen -= subLen
				index += subLen
			} else {
				index += subLen
				break
			}
		}
	}
	return index
}
