package findNthDigit

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	i := 1
	base := 1
	count := i * base * 9
	for n > count {
		n -= count
		i++
		base *= 10
		count = i * base * 9
	}
	num := (n-1)/i + base
	for j := 0; j < i-(n-1)%i-1; j++ {
		num /= 10
	}
	return num % 10
}
