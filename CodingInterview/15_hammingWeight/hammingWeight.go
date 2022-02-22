package hammingWeight

// // 判断奇偶用取余
// func hammingWeight(num uint32) int {
// 	count := 0
// 	for ; num != 0; num >>= 1 {
// 		count += int(num % 2)
// 	}
// 	return count
// }

// 使用和uint32(1)取或是否等于0也可以判断奇偶，复杂度比取余低
func hammingWeight(num uint32) int {
	count := 0
	one := uint32(1)
	for ; num != 0; num >>= 1 {
		if one&num != 0 {
			count++
		}
	}
	return count
}
