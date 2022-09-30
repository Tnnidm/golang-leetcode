package goodIndices

func goodIndices(nums []int, k int) []int {
	postIncrease := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 && nums[i] <= nums[i+1] {
			postIncrease[i] = postIncrease[i+1] + 1
		} else {
			postIncrease[i] = 1
		}
	}
	preDecrease := 1
	for i := 1; i < k; i++ {
		if nums[i-1] >= nums[i] {
			preDecrease++
		} else {
			preDecrease = 1
		}
	}
	result := []int{}
	for i := k; i < len(nums)-k; i++ {
		if preDecrease >= k && postIncrease[i+1] >= k {
			result = append(result, i)
		}
		if nums[i-1] >= nums[i] {
			preDecrease++
		} else {
			preDecrease = 1
		}
	}
	return result
}
