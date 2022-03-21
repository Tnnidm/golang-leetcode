package myPow

// 迭代写法
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var negFlag bool
	if n < 0 {
		n = -n
		negFlag = true
	}
	result := 1.0
	for n != 0 {
		if n&1 != 0 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	if negFlag {
		result = 1 / result
	}
	return result
}
