package longestPalindrome

func longestPalindrome(s string) int {
	charMap := make([]int, 52)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			charMap[s[i]-'a'] += 1
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			charMap[s[i]-'A'+26] += 1
		}
	}
	singleNum := 0
	sum := 0
	for i := 0; i < 52; i++ {
		sum += charMap[i]
		if charMap[i]%2 != 0 {
			singleNum++
		}
	}
	if singleNum != 0 {
		sum -= (singleNum - 1)
	}
	return sum
}
