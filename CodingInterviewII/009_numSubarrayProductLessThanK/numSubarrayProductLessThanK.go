package numSubarrayProductLessThanK

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	var result int
	product := 1
	for right < len(nums) {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		if left <= right {
			result += right - left + 1
		}
		right++
	}
	return result
}
