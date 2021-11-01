package findDisappearedNumbers

// // 这种方法简单直接但是复杂度比较高，需要O(n)的空间，能不能只需要O(1)的空间呢
// func findDisappearedNumbers(nums []int) []int {
// 	numsLen := len(nums)
// 	numIndex := make([]int, numsLen)
// 	for i := 0; i < numsLen; i++ {
// 		numIndex[nums[i]-1]++
// 	}
// 	result := []int{}
// 	for i := 0; i < numsLen; i++ {
// 		if numIndex[i] == 0 {
// 			result = append(result, i+1)
// 		}
// 	}
// 	return result
// }

// 因为这个表的长度和数值都是1到n所以可以拿这个表本身做自身哈希
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
