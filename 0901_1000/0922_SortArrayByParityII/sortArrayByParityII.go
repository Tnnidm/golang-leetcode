package sortArrayByParityII

func sortArrayByParityII(nums []int) []int {
	for i, j := 0, 0; i < len(nums) && j < len(nums); i, j = i+1, j+1 {
		for !(i%2 == 0 && nums[i]%2 == 1) && i < len(nums)-1 {
			i++
		}
		for !(j%2 == 1 && nums[j]%2 == 0) && j < len(nums)-1 {
			j++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
