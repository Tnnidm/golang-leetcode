package isThree

func isThree(n int) bool {
	if n < 4 {
		return false
	}
	divideNum := 2
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divideNum++
		}
	}
	return divideNum == 3
}
