package findSubstring

// func findSubstring(s string, words []string) []int {
// 	wordCount := []int{}
// 	wordIndex := map[string]int{}
// 	index := -1
// 	for _, word := range words {
// 		if i, ok := wordIndex[word]; ok {
// 			wordCount[i]++
// 		} else {
// 			wordCount = append(wordCount, 1)
// 			index++
// 			wordIndex[word] = index
// 		}
// 	}
// 	wordCountLen := len(wordCount)
// 	wordNum := len(words)
// 	wordLen := len(words[0])
// 	result := []int{}
// 	for i := 0; i <= len(s)-wordNum*wordLen; i++ {
// 		wordCountTemp := make([]int, wordCountLen)
// 		completed := true
// 		for j := 0; j < wordNum; j++ {
// 			if index, ok := wordIndex[s[i+j*wordLen:i+(j+1)*wordLen]]; ok {
// 				wordCountTemp[index]++
// 			} else {
// 				completed = false
// 				break
// 			}
// 		}
// 		if completed && isEqual(wordCountTemp, wordCount, wordCountLen) {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

// func isEqual(list1, list2 []int, length int) bool {
// 	for i := 0; i < length; i++ {
// 		if list1[i] != list2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func findSubstring(s string, words []string) []int {
	wordNum := len(words)
	wordLen := len(words[0])
	wordCount := make(map[string]int, wordNum)
	wordsTotoalLength := wordLen * wordNum
	sLen := len(s)
	if wordsTotoalLength > sLen {
		return []int{}
	}
	result := []int{}
	for _, word := range words {
		wordCount[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0 //表示已经符合的数量
		left, right := i, i
		sstring := map[string]int{}
		for right+wordLen <= sLen {
			word := s[right : right+wordLen]
			right += wordLen
			if _, ok := wordCount[word]; !ok { //不存在
				left = right
				sstring = map[string]int{} //清空map
				count = 0
			} else {
				sstring[word]++
				count++
				for sstring[word] > wordCount[word] {
					temp := s[left : left+wordLen]
					sstring[temp]--
					count--
					left += wordLen
				}
				if count == wordNum {
					result = append(result, left)
				}
			}
		}

	}
	return result
}
