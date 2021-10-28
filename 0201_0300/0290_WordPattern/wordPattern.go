package wordPattern

import "strings"

// import "strings"

func wordPattern(pattern string, s string) bool {
	// 可以直接采用
	splitString := strings.Split(s, " ") //比自己写的内存消耗小点
	// splitString := splitStr(s)

	splitStringLen := len(splitString)
	patternLen := len(pattern)
	if splitStringLen != patternLen {
		return false
	} else {
		patternMap := make(map[byte]string)
		strMap := make(map[string]byte)
		for i := 0; i < patternLen; i++ {
			if word, ok := patternMap[pattern[i]]; ok {
				if word != splitString[i] {
					return false
				}
			} else {
				patternMap[pattern[i]] = splitString[i]
			}
			if apattern, ok := strMap[splitString[i]]; ok {
				if apattern != pattern[i] {
					return false
				}
			} else {
				strMap[splitString[i]] = pattern[i]
			}
		}
		return true
	}
}

func splitStr(s string) []string {
	splitResult := []string{}
	begin := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i] == ' ' {
			splitResult = append(splitResult, s[begin:i])
			begin = i + 1
		} else if i == sLen-1 {
			splitResult = append(splitResult, s[begin:sLen])
		}
	}
	return splitResult
}
