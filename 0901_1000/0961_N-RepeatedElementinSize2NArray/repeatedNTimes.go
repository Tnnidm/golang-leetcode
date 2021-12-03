package repeatedNTimes

func repeatedNTimes(nums []int) int {
	numsLen := len(nums)
	amap := make(map[int]bool, numsLen/2)
	for i := 0; i < numsLen; i++ {
		if amap[nums[i]] == false {
			amap[nums[i]] = true
		} else {
			return nums[i]
		}
	}
	return 0
}
