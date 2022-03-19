package singleNumbers

func twoSingleNumbers(nums []int) []int {
	numsLen := len(nums)
	temp := 0
	for i := 0; i < numsLen; i++ {
		temp ^= nums[i]
	}
	k := 1
	for temp&k == 0 {
		k <<= 1
	}
	ans := []int{0, 0}
	for i := 0; i < numsLen; i++ {
		if nums[i]&k == 0 {
			ans[0] ^= nums[i]
		} else {
			ans[1] ^= nums[i]
		}
	}
	return ans
}
