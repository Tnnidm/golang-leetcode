package numberOfMatches

func numberOfMatches(n int) int {
	if n <= 2 {
		return n - 1
	}
	if n%2 == 0 {
		return numberOfMatches(n/2) + n/2
	} else {
		return numberOfMatches((n-1)/2+1) + (n-1)/2
	}
}
