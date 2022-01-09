package sumBase

func sumBase(n int, k int) int {
	result := 0
	for n != 0 {
		result += n % k
		n = n / k
	}
	return result
}
