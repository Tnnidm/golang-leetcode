package groupAnagrams

// 对字符串排序法
// func groupAnagrams(strs []string) [][]string {
// 	groupIndex := map[string]int{}
// 	result := [][]string{}
// 	for _, str := range strs {
// 		temp := stringSort(str)
// 		if index, ok := groupIndex[temp]; ok {
// 			result[index] = append(result[index], str)
// 		} else {
// 			groupIndex[temp] = len(result)
// 			result = append(result, []string{str})
// 		}
// 	}
// 	return result
// }

// func stringSort(str string) string {
// 	chars := []byte(str)
// 	sort.Slice(chars, func(i, j int) bool {
// 		return chars[i] <= chars[j]
// 	})
// 	return string(chars)
// }

// 对字符串统计法
func groupAnagrams(strs []string) [][]string {
	groupIndex := map[[26]int]int{}
	result := [][]string{}
	for _, str := range strs {
		temp := stringStat(str)
		if index, ok := groupIndex[temp]; ok {
			result[index] = append(result[index], str)
		} else {
			groupIndex[temp] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}

func stringStat(str string) [26]int {
	var chars [26]int
	for _, char := range str {
		chars[char-'a']++
	}
	return chars
}
