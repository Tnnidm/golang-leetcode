package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	groupMap := map[[26]int]int{}
	for _, str := range strs {
		var charCount [26]int
		for _, char := range str {
			charCount[char-'a']++
		}
		if index, ok := groupMap[charCount]; ok {
			result[index] = append(result[index], str)
		} else {
			groupMap[charCount] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}
