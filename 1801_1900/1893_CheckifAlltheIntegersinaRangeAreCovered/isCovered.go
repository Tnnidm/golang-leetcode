package isCovered

func isCovered(ranges [][]int, left int, right int) bool {
	isCovered := make([]bool, 51)
	for i := range ranges {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			isCovered[j] = true
		}
	}
	for i := left; i <= right; i++ {
		if isCovered[i] == false {
			return false
		}
	}
	return true
}
