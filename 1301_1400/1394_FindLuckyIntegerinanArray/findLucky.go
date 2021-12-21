package findLucky

func findLucky(arr []int) int {
	hash := make([]int, 501)
	for _, d := range arr {
		hash[d]++
	}

	for i := len(hash) - 1; i > 0; i-- {
		if i == hash[i] {
			return i
		}
	}
	return -1

}
