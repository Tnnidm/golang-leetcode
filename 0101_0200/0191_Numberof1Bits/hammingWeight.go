package hammingWeight

func hammingWeight(num uint32) int {
	result := 0
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			result += 1
		}
		num /= 2
	}
	return result
}

// // 内存消耗更低的方案
// func hammingWeight(num uint32) int {
// 	var result int
// 	for ; num > 0; num = num & (num - 1) {
// 		result++
// 	}
// 	return result
// }
