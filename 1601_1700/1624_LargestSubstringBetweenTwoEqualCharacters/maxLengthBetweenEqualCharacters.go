package maxLengthBetweenEqualCharacters

func maxLengthBetweenEqualCharacters(s string) int {
	maxDistance := -1
	charLastIndex := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if lastIndex, ok := charLastIndex[s[i]]; ok {
			if i-lastIndex-1 > maxDistance {
				maxDistance = i - lastIndex - 1
			}
		} else {
			charLastIndex[s[i]] = i
		}
	}
	return maxDistance
}
