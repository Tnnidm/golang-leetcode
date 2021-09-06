package removeDuplicates

// My method
// 执行用时：8 ms, 在所有 Go 提交中击败了87.05%的用户
// 内存消耗：4.3 MB, 在所有 Go 提交中击败了99.94%的用户
func removeDuplicates(nums []int) int {
	var counter int = 0
	var index int = 0
	for counter != len(nums) {
		for index = counter; ; index++ {
			if index == len(nums)-1 {
				break
			} else if nums[index] != nums[index+1] {
				break
			}
		}
		counter = counter + 1
		nums = append(nums[:counter], nums[index+1:]...)
	}
	return counter
}
