package divide

import "math"

func divide(a int, b int) int {
	if a == math.MinInt32 {
		switch b {
		case -1:
			return math.MaxInt32
		case 1:
			return math.MinInt32
		case math.MinInt32:
			return 1
		}
	}
	if a == 0 || b == math.MinInt32 {
		return 0
	}
	var falseFlag bool
	if a < 0 {
		a = -a
		falseFlag = !falseFlag
	}
	if b < 0 {
		b = -b
		falseFlag = !falseFlag
	}
	temp, count := b, 1
	for temp<<1 <= a {
		temp <<= 1
		count <<= 1
	}
	result := 0
	for a >= b {
		if a >= temp {
			result += count
			a -= temp
		}
		temp >>= 1
		count >>= 1
	}
	if falseFlag {
		return -result
	}
	return result
}
