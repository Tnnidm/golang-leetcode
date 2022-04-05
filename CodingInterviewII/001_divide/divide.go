package divide

import "math"

func divide(a int, b int) int {
	var negFlag bool
	if a < 0 {
		a = -a
		negFlag = !negFlag
	}
	if b < 0 {
		b = -b
		negFlag = !negFlag
	}
	if a < b {
		return 0
	}
	result := 0
	count := 1
	var base int
	for base = b; base<<1 <= a; base <<= 1 {
		count <<= 1
		if count > math.MaxInt32 {
			if negFlag {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	result += count
	a -= base
	for a != 0 {
		for base > a {
			base >>= 1
			count >>= 1
		}
		a -= base
		result += count
	}
	if negFlag {
		return -result
	}
	return result
}
