package hammingWeight

// // 判断奇偶用取余
// func hammingWeight(num uint32) int {
// 	count := 0
// 	for ; num != 0; num >>= 1 {
// 		count += int(num % 2)
// 	}
// 	return count
// }

// // 使用和uint32(1)取或是否等于0也可以判断奇偶，复杂度比取余低
// func hammingWeight(num uint32) int {
// 	count := 0
// 	one := uint32(1)
// 	for ; num != 0; num >>= 1 {
// 		if one&num != 0 {
// 			count++
// 		}
// 	}
// 	return count
// }

// // 也可以反过来增加uint32(1)的位数
// func hammingWeight(num uint32) int {
// 	count := 0
// 	one := uint32(1)
// 	for ; one != 0; one <<= 1 {
// 		if one&num != 0 {
// 			count++
// 		}
// 	}
// 	return count
// }

// 因为n和n-1做与操作后，n的最低位的1会变成0，那么在n整个变成0之前，能做几次这种操作就是有几个1
func hammingWeight(num uint32) int {
	count := 0
	for ; num != 0; num = num & (num - 1) {
		count++
	}
	return count
}
