package pivotIndex

// func pivotIndex(nums []int) int {
// 	numsLen := len(nums)
// 	LeftSum := 0
// 	RightSum := 0
// 	for i := 1; i < numsLen; i++ {
// 		RightSum += nums[i]
// 	}
// 	nums = append(nums, 0)
// 	for i := 0; i < numsLen; i++ {
// 		if LeftSum == RightSum {
// 			return i
// 		} else {
// 			LeftSum += nums[i]
// 			RightSum -= nums[i+1]
// 		}
// 	}
// 	return -1
// }

func pivotIndex(nums []int) int {
	sum := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		sum += nums[i]
	}
	leftSum := 0
	for i := 0; i < numsLen; i++ {
		if 2*leftSum+nums[i] == sum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
