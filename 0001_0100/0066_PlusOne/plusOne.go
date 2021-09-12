package plusOne

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	digits[digitsLen-1] += 1
	for i := digitsLen - 1; i >= 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			if i != 0 {
				digits[i-1] += 1
			} else {
				digits = append([]int{1}, digits...)
				break
			}
		} else {
			break
		}
	}
	return digits
}
