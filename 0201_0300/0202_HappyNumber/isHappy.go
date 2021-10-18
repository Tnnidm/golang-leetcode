package isHappy

func isHappy(n int) bool {
	nMap := map[int]struct{}{}
	for counter := 0; counter < 1000; counter++ {
		if n == 1 {
			return true
		} else {
			nMap[n] = struct{}{}
			quotient, remainder := n, 0
			n = 0
			for quotient != 0 {
				remainder = quotient % 10
				n = n + remainder*remainder
				quotient = quotient / 10
			}
			if _, ok := nMap[n]; ok {
				return false
			}
		}
	}
	return false
}
