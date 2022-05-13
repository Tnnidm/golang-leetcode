package minFlipsMonoIncr

func minFlipsMonoIncr(s string) int {
	endWithOne, endWithZero := 0, 0
	for _, num := range s {
		if num == '0' {
			endWithOne++
		} else {
			endWithOne = min(endWithOne, endWithZero)
			endWithZero++
		}
	}
	return min(endWithOne, endWithZero)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
