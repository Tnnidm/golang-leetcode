package maxNumberOfBalloons

func maxNumberOfBalloons(text string) int {
	charMap := make([]int, 26)
	for _, v := range text {
		charMap[v-'a']++
	}
	minNum := 10000
	if charMap['b'-'a'] < minNum {
		minNum = charMap['b'-'a']
	}
	if charMap['a'-'a'] < minNum {
		minNum = charMap['a'-'a']
	}
	if charMap['l'-'a']/2 < minNum {
		minNum = charMap['l'-'a'] / 2
	}
	if charMap['o'-'a']/2 < minNum {
		minNum = charMap['o'-'a'] / 2
	}
	if charMap['n'-'a'] < minNum {
		minNum = charMap['n'-'a']
	}
	return minNum
}
