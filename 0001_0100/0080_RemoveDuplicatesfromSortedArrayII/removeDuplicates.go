package removeDuplicates

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	slow := 1
	count := 1
	for fast := 1; fast < numsLen; fast++ {
		if nums[fast] == nums[fast-1] {
			count++
			if count <= 2 {
				nums[slow] = nums[fast]
				slow++
			}
		} else {
			count = 1
			nums[slow] = nums[fast]
			slow++
		}
		// fmt.Println(nums, count, fast, slow)
	}
	nums = nums[:slow]
	return slow
}
