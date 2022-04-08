package minSubArrayLen

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := nums[0]
	numsLen := len(nums)
	minLen := 100001
	for {
		if sum < target {
			if right < numsLen-1 {
				right++
				sum += nums[right]
			} else {
				break
			}
		} else {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			if left < right {
				sum -= nums[left]
				left++
			} else {
				if left != numsLen-1 {
					left++
					right++
					sum = nums[left]
				} else {
					break
				}
			}
		}
	}
	if minLen == 100001 {
		return 0
	}
	return minLen
}
