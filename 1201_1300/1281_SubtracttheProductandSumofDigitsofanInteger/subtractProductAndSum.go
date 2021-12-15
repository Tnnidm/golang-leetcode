package subtractProductAndSum

func subtractProductAndSum(n int) int {
	times := 1
	sum := 0
	for n != 0 {
		times = times * (n % 10)
		sum = sum + n%10
		n = n / 10
	}
	return times - sum
}
