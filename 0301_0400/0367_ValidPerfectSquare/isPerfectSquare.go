package isPerfectSquare

func isPerfectSquare(num int) bool {
	for i := 1; ; i++ {
		if i*i == num {
			return true
		} else if i*i > num {
			return false
		}
	}
}
