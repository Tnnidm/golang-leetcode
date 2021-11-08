package checkPerfectNumber

import "math"

func checkPerfectNumber(num int) bool {
	if num < 4 {
		return num == 1
	} else {
		sum := 1
		numSqrt := int(math.Sqrt(float64(num)))
		for i := 2; i <= numSqrt; i++ {
			if num%i == 0 {
				if i != numSqrt {
					sum += (i + num/i)
				} else {
					sum += i
				}
			}
		}
		return sum == num
	}
}
