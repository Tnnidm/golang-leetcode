package isPowerOfTwo

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	} else {
		for n != 1 {
			if n%2 == 1 {
				return false
			}
			n = n / 2
		}
		return true
	}
}
