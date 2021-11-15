package hasAlternatingBits

func hasAlternatingBits(n int) bool {
	bit := 1 - n%2
	for n != 0 {
		if n%2 != bit {
			bit = n % 2
			n = n / 2
		} else {
			return false
		}
	}
	return true
}

// // 另一种解法，消耗更小内存
// func hasAlternatingBits(n int) bool {
// 	num := n ^ (n >> 1) + 1
// 	return num&(num-1) == 0
// }
