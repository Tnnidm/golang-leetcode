package canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	charMap := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		charMap[magazine[i]-'a'] += 1
	}
	for i := 0; i < len(ransomNote); i++ {
		charMap[ransomNote[i]-'a'] -= 1
		if charMap[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
