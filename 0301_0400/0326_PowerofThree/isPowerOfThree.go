package isPowerOfThree

// func isPowerOfThree(n int) bool {
// 	if n < 1 {
// 		return false
// 	} else {
// 		for n != 1 {
// 			if n%3 != 0 {
// 				return false
// 			} else {
// 				n = n / 3
// 			}
// 		}
// 		return true
// 	}
// }

// 有更简洁的写法
func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}
