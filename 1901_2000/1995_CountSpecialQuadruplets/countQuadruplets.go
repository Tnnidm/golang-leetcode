package countQuadruplets

func countQuadruplets(nums []int) int {
	numsLen := len(nums)
	result := 0
	for i := 0; i < numsLen-3; i++ {
		for j := i + 1; j < numsLen-2; j++ {
			for k := j + 1; k < numsLen-1; k++ {
				for m := k + 1; m < numsLen; m++ {
					if nums[i]+nums[j]+nums[k] == nums[m] {
						result++
					}
				}
			}
		}
	}
	return result
}
