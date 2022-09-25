package maxSlidingWindow

func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	dequeLastIndex := -1
	deque := []int{}
	result := make([]int, numsLen-k+1)
	for right := 0; right < k; right++ {
		for dequeLastIndex >= 0 && deque[dequeLastIndex] < nums[right] {
			deque = deque[:dequeLastIndex]
			dequeLastIndex--
		}
		deque = append(deque, nums[right])
		dequeLastIndex++
	}
	result[0] = deque[0]
	for left, right := 0, k; right < numsLen; left, right = left+1, right+1 {
		if nums[left] == deque[0] {
			deque = deque[1:]
			dequeLastIndex--
		}
		for dequeLastIndex >= 0 && deque[dequeLastIndex] < nums[right] {
			deque = deque[:dequeLastIndex]
			dequeLastIndex--
		}
		deque = append(deque, nums[right])
		dequeLastIndex++
		result[left+1] = deque[0]
	}
	return result
}
