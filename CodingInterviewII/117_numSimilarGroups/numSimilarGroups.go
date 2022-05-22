package numSimilarGroups

func numSimilarGroups(strs []string) int {
	strLen := len(strs)
	fathers := make([]int, strLen)
	for i := 0; i < strLen; i++ {
		fathers[i] = i
	}
	result := strLen
	for i := 0; i < strLen; i++ {
		for j := i + 1; j < strLen; j++ {
			if isLike(strs[i], strs[j]) {
				fatherOfi := findFather(fathers, i)
				fatherOfj := findFather(fathers, j)
				if fatherOfi != fatherOfj {
					fathers[fatherOfi] = fatherOfj
					result--
				}
			}
		}
	}
	return result
}

func isLike(s1, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
		if count > 2 {
			return false
		}
	}
	return true
}

func findFather(fathers []int, i int) int {
	if i != fathers[i] {
		fathers[i] = findFather(fathers, fathers[i])
	}
	return fathers[i]
}
