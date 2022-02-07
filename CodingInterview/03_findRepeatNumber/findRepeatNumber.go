package findRepeatNumber

// // 解法1: 哈希表，时间复杂度O(n), 空间复杂度O(n)
// func findRepeatNumber(nums []int) int {
// 	N := len(nums)
// 	hashMap := make(map[int]bool, N)
// 	for _, num := range nums {
// 		if hashMap[num] {
// 			return num
// 		}
// 		hashMap[num] = true
// 	}
// 	return 0
// }

// // 解法2: 数字实现简单哈希表，时间复杂度O(n)，空间复杂度O(n)。但因为用下标代替了哈希表键，实际效果会比哈希表好得多
// func findRepeatNumber(nums []int) int {
// 	N := len(nums)
// 	hashMap := make([]bool, N)
// 	for _, num := range nums {
// 		if hashMap[num] {
// 			return num
// 		}
// 		hashMap[num] = true
// 	}
// 	return 0
// }

// // 解法3: 排序后遍历找重复的数字，时间复杂度O(nlogn)，空间复杂度O(1)
// func findRepeatNumber(nums []int) int {
// 	sort.Ints(nums)
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			return nums[i]
// 		}
// 	}
// 	return 0
// }

// 解法4: 利用第i个元素和第nums[i]个元素交换，一定会使nums[i]出现在第nums[i]位置上这个属性
// 时间复杂度O(n)，空间复杂度O(1)
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return 0
}
