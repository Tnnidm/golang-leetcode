package mySqrt

import "math"

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x < 4 {
		return 1
	} else {
		upperBound := min(math.MaxUint16, x)
		lowerBound := 2
		half := 0
		half_square := 0
		for {
			if upperBound-lowerBound == 1 {
				return lowerBound
			}
			half = (upperBound + lowerBound) / 2
			half_square = half * half
			if x < half_square {
				upperBound = half
			} else if x > half_square {
				lowerBound = half
			} else {
				return half
			}

		}
	}

}
