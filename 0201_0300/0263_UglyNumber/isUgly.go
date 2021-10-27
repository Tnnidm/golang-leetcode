package isUgly

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n != 1 {
		notuglyFlag := true
		if n%2 == 0 {
			notuglyFlag = false
			n = n / 2
		}
		if n%3 == 0 {
			notuglyFlag = false
			n = n / 3
		}
		if n%5 == 0 {
			notuglyFlag = false
			n = n / 5
		}
		if notuglyFlag {
			return false
		}
	}
	return true
}
