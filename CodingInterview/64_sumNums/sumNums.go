package sumNums

// 递归法，利用逻辑运算符的短路性质
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(i int) bool {
		ans += i
		return i > 1 && sum(i-1)
	}
	sum(n)
	return ans
}

// // 快速乘，也叫俄罗斯农民乘法
// func sumNums(n int) int {
// 	ans, A, B := 0, n, n+1

// 	addGreatZero := func() bool {
// 		ans += A
// 		return ans > 0
// 	}

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1

// 	_ = ((B & 1) > 0) && addGreatZero()
// 	A <<= 1
// 	B >>= 1
// 	return ans >> 1
// }

// func quickMulti(A, B int) int {
// 	ans := 0
// 	for ; B != 0; B >>= 1 {
// 		if B&1 != 0 {
// 			ans += A
// 		}
// 		A <<= 1
// 	}
// 	return ans
// }
