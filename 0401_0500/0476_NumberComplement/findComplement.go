package findComplement

func findComplement(num int) int {
	result := 0
	count := 0
	for num != 0 {
		result += (1 - num%2) << count
		num /= 2
		count++
	}
	return result
}
