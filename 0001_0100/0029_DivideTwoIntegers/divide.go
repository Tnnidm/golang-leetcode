package divide

import "math"

func divide(dividend int, divisor int) int {
	var dividendNegFlag, divisorNegFlag bool
	if dividend < 0 {
		dividend = -dividend
		dividendNegFlag = true
	}
	if divisor < 0 {
		divisor = -divisor
		divisorNegFlag = true
	}
	if dividend < divisor {
		return 0
	}
	result := 0
	count := 1
	var base int
	for base = divisor; base<<1 <= dividend; base <<= 1 {
		count <<= 1
		if count > math.MaxInt32 {
			if dividendNegFlag != divisorNegFlag {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	result += count
	dividend -= base
	for dividend != 0 {
		for base > dividend {
			base >>= 1
			count >>= 1
		}
		dividend -= base
		result += count
	}
	if dividendNegFlag != divisorNegFlag {
		return -result
	}
	return result
}
