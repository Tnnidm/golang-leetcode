package sumZero

func sumZero(n int) []int {
	ans := make([]int, n)
	for i := 0; i < n/2; i++ {
		ans[2*i] = i + 1
		ans[2*i+1] = -i - 1
	}
	return ans
}
