package countTriples

func countTriples(n int) int {
	var iSquare, jSquare, result int
	for i := 1; i <= n; i++ {
		iSquare = i * i
		for j := 1; j <= n; j++ {
			jSquare = j * j
			for k := i; k <= n; k++ {
				if iSquare+jSquare == k*k {
					result++
				}
			}
		}
	}
	return result
}
