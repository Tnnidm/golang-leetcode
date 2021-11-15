package findLengthOfLCIS

func findLengthOfLCIS(nums []int) int {
	LCIS := 1
	numsLen := len(nums)
	if numsLen == 1 {
		return LCIS
	} else {
		CIS := 1
		for i := 1; i < numsLen; i++ {
			if nums[i] > nums[i-1] {
				CIS++
				if CIS > LCIS {
					LCIS = CIS
				}
			} else {
				CIS = 1
			}
		}
		return LCIS
	}
}
