package findNumbers

func findNumbers(nums []int) int {
	result := 0
	bit := 0
	for _, num := range nums {
		for bit = 0; num != 0; num = num / 10 {
			bit++
		}
		if bit%2 == 0 {
			result++
		}
	}
	return result
}
