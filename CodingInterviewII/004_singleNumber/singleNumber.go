package singleNumber

// 如果是先遍历num然后按位统计total，会有溢出的风险
// 这种先遍历位，然后遍历num的方法更安全
func singleNumber(nums []int) int {
	result := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		result |= total % 3 << i
	}
	return int(result)
}

// // 我自己的思路也是没有问题的
// // 就是64位cpu需要纠正一下整数位数
// func singleNumber(nums []int) int {
// 	bits := make([]int, 32)
// 	for _, num := range nums {
// 		for i := 0; i < 32; i++ {
// 			bits[i] += (num & (1 << i)) >> i
// 		}
// 	}
// 	result := 0
// 	for i := 0; i < 32; i++ {
// 		result |= bits[i] % 3 << i
// 	}
// 	return int(int32(result)) // 这里非常重要！！！！
// }
