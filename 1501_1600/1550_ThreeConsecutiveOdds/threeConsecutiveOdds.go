package threeConsecutiveOdds

func threeConsecutiveOdds(arr []int) bool {
	continuousOddCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			continuousOddCount++
			if continuousOddCount == 3 {
				return true
			}
		} else {
			continuousOddCount = 0
		}
	}
	return false
}
