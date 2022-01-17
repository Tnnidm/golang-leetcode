package canBeIncreasing

func canBeIncreasing(nums []int) bool {
	decreaseTimes := 0
	index := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			decreaseTimes++
			index = i - 1
			if decreaseTimes > 1 {
				return false
			}
		}
	}
	if decreaseTimes == 0 {
		return true
	} else {
		return (index == 0 || nums[index-1] < nums[index+1]) || (index+2 == len(nums) || nums[index] < nums[index+2])
	}
}
