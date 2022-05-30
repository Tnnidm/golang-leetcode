package sortColors

// // 计数排序
// func sortColors(nums []int) {
// 	var stat [3]int
// 	for _, num := range nums {
// 		stat[num]++
// 	}
// 	var index int
// 	for num := 0; num <= 2; num++ {
// 		for i := 0; i < stat[num]; i++ {
// 			nums[index] = num
// 			index++
// 		}
// 	}
// }

// 原地排序
func sortColors(nums []int) {
	numsLen := len(nums)
	ptr0 := 0
	for i := 0; i < numsLen; i++ {
		if nums[i] == 0 {
			nums[i], nums[ptr0] = nums[ptr0], nums[i]
			ptr0++
		}
	}
	ptr2 := numsLen - 1
	for i := numsLen - 1; i >= ptr0; i-- {
		if nums[i] == 2 {
			nums[i], nums[ptr2] = nums[ptr2], nums[i]
			ptr2--
		}
	}
}
