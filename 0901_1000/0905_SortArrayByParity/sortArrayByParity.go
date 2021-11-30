package sortArrayByParity

// func sortArrayByParity(nums []int) []int {
// 	for i := range nums {
// 		if nums[i]%2 == 1 {
// 			nums[i] += 5000
// 		}
// 	}
// 	sort.Ints(nums)
// 	for i := range nums {
// 		if nums[len(nums)-1-i]%2 == 1 {
// 			nums[len(nums)-1-i] -= 5000
// 		} else {
// 			break
// 		}
// 	}
// 	return nums
// }

func sortArrayByParity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		for nums[i]%2 == 0 && i < j {
			i++
		}
		for nums[j]%2 == 1 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
