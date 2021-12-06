package prefixesDivBy5

func prefixesDivBy5(nums []int) []bool {
	numsLen := len(nums)
	result := make([]bool, numsLen)
	num := 0
	for i := 0; i < numsLen; i++ {
		num = (num*2 + nums[i]) % 5
		result[i] = num == 0
	}
	return result
}
