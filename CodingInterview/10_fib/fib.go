package fib

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		result := [3]int{0, 1, 1}
		for ; n > 1; n-- {
			result[2] = (result[1] + result[0]) % 1000000007
			result[0] = result[1]
			result[1] = result[2]
		}
		return result[2]
	}
}
