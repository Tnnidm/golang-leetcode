package myPow

import "math"

// 迭代写法
func myPow(x float64, n int) float64 {
	if x == 0 {
		if n < 0 {
			return math.Inf(0)
		} else if n == 0 {
			return 1
		} else {
			return 0
		}
	} else if x == 1 {
		return 1
	} else {
		negFlag := false
		if n < 0 {
			n = -n
			negFlag = true
		}
		base := x
		result := 1.0000
		for ; n != 0; n >>= 1 {
			if n&1 != 0 {
				result *= base
			}
			base *= base
		}
		if negFlag {
			result = 1 / result
		}
		return result
	}
}

// // 递归写法
// func myPow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1.0
// 	} else if n < 0 {
// 		return 1.0 / myPow(x, -n)
// 	}
// 	temp := myPow(x, n/2)
// 	if n%2 == 0 {
// 		return temp * temp
// 	}
// 	return x * temp * temp
// }
