package maxSlidingWindow

func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{}
	}
	result := make([]int, numsLen-k+1)
	areaMax := nums[0]
	maxIndex := 0
	for i := 0; i < k; i++ {
		if nums[i] >= areaMax {
			areaMax = nums[i]
			maxIndex = i
		}
	}
	result[0] = areaMax
	for i := k; i < numsLen; i++ {
		if nums[i] >= areaMax {
			areaMax = nums[i]
			maxIndex = i
		} else if i-k == maxIndex {
			maxIndex = i - k + 1
			areaMax = nums[maxIndex]
			for j := i - k + 2; j <= i; j++ {
				if nums[j] >= areaMax {
					areaMax = nums[j]
					maxIndex = j
				}
			}
		}
		result[i-k+1] = areaMax
	}
	return result
}
