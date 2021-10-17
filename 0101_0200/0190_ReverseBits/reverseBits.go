package reverseBits

func reverseBits(num uint32) uint32 {
	var bit uint32
	var result uint32
	for i := 0; i < 32; i++ {
		bit = num % 2
		result = 2*result + bit
		num = num / 2
	}
	return result
}

// // 消耗更小内存的写法
// func reverseBits(num uint32) uint32 {
//     var ans uint32
// 	for i := 0; i < 32; i++ {
// 		ans <<= 1
// 		ans |= num >> i & 1
// 	}
// 	return ans
// }
