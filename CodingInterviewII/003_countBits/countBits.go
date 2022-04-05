package countBits

// // 方法1:用i&(i-1)暴力计算
// func countBits(n int) []int {
// 	result := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		for temp := i; temp != 0; temp &= (temp - 1) {
// 			result[i]++
// 		}
// 	}
// 	return result
// }

// // 方法2: 因为i中1的数量始终是在i&(i-1)中1的数量后计算的，且i中1的数量就比&(i-1)中1的数量多1
// func countBits(n int) []int {
// 	result := make([]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		result[i] = result[i&(i-1)]+1
// 	}
// 	return result
// }

// 方法2: 因为i中1的数量始终是在i/2中1的数量后计算的，且只有i为奇数的时候，i中1的数量才比i/2中1的数量多1
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
