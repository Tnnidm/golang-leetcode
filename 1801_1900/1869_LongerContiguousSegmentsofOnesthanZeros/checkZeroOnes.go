package checkZeroOnes

func checkZeroOnes(s string) bool {
	continuousOneNum, maxContinuousOneNum := 0, 0
	continuousZeroNum, maxContinuousZeroNum := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continuousOneNum = 0
			continuousZeroNum++
			if continuousZeroNum > maxContinuousZeroNum {
				maxContinuousZeroNum = continuousZeroNum
			}
		} else {
			continuousZeroNum = 0
			continuousOneNum++
			if continuousOneNum > maxContinuousOneNum {
				maxContinuousOneNum = continuousOneNum
			}
		}
	}
	return maxContinuousOneNum > maxContinuousZeroNum
}
