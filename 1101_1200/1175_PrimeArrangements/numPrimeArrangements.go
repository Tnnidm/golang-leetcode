package numPrimeArrangements

func numPrimeArrangements(n int) int {
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	base := 1000000007
	num := 0
	for ; num < 25 && prime[num] <= n; num++ {
	}
	ans := 1
	if num > 1 {
		for i := 2; i <= num; i++ {
			ans = (ans * i) % base
		}
	}
	if n-num > 1 {
		for i := 2; i <= n-num; i++ {
			ans = (ans * i) % base
		}
	}
	return ans
}
