package numberOfWays

func numberOfWays(startPos int, endPos int, k int) int {
	necessary := abs(endPos - startPos)
	if necessary > k || (k-necessary)%2 == 1 {
		return 0
	}
	temp := (k - necessary) >> 1
	memory := map[[2]int]int{}
	var combination func(n, m int) int
	combination = func(n, m int) int {
		if n == 0 || n == m {
			memory[[2]int{n, m}] = 1
			return 1
		}
		if memory[[2]int{n, m}] == 0 {
			memory[[2]int{n, m}] = (combination(n-1, m-1) + combination(n, m-1)) % 1000000007
		}
		return memory[[2]int{n, m}]
	}
	return combination(temp, k)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
