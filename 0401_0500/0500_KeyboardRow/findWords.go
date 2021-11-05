package findWords

import "strings"

func findWords(words []string) []string {
	var result []string
	for _, w := range words {
		wLower := strings.ToLower(w)
		for _, kbrow := range []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"} {
			flag := true
			for _, char := range wLower {
				if !strings.ContainsRune(kbrow, char) {
					flag = false
					break
				}
			}
			if flag {
				result = append(result, w)
				break
			}
		}
	}
	return result
}
