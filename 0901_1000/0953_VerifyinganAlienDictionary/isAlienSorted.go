package isAlienSorted

func isAlienSorted(words []string, order string) bool {
	alienMap := make([]int, len(order))
	for i := range order {
		alienMap[order[i]-'a'] = i
	}
	N := len(words)
	result := true
	var equalFlag bool
	var word1Len, word2Len int
	for i := 0; i < N-1 && result; i++ {
		equalFlag = true
		word1Len = len(words[i])
		word2Len = len(words[i+1])
		for j := 0; j < min(word1Len, word2Len); j++ {
			if alienMap[words[i+1][j]-'a'] > alienMap[words[i][j]-'a'] {
				equalFlag = false
				break
			} else if alienMap[words[i+1][j]-'a'] < alienMap[words[i][j]-'a'] {
				result = false
				equalFlag = false
				break
			}
		}
		if equalFlag {
			result = word1Len <= word2Len
		}
	}
	return result
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
