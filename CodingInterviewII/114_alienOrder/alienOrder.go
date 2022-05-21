package alienOrder

import "strings"

func alienOrder(words []string) string {
	graph := [26][]byte{}
	inDegree := [26]int{}
	exist := [26]bool{}
	// 处理第一个word的
	lastWordLen := len(words[0])
	for j := 0; j < lastWordLen; j++ {
		exist[words[0][j]-'a'] = true
	}
	var found bool
	for i := 1; i < len(words); i++ {
		wordLen := len(words[i])
		for j := 0; j < wordLen; j++ {
			exist[words[i][j]-'a'] = true
		}
		found = false
		for j := 0; j < min(lastWordLen, wordLen); j++ {
			if words[i-1][j] != words[i][j] {
				found = true
				graph[words[i-1][j]-'a'] = append(graph[words[i-1][j]-'a'], words[i][j])
				inDegree[words[i][j]-'a']++
				break
			}
		}
		if !found && lastWordLen > wordLen {
			return ""
		}
		lastWordLen = wordLen
	}
	result := strings.Builder{}
	queue := []int{}
	for i := 0; i < 26; i++ {
		if exist[i] && inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		zeroInNode := queue[0]
		queue = queue[1:]
		result.WriteByte('a' + uint8(zeroInNode))
		for _, char := range graph[zeroInNode] {
			inDegree[char-'a']--
			if inDegree[char-'a'] == 0 {
				queue = append(queue, int(char-'a'))
			}
		}
	}
	for i := 0; i < 26; i++ {
		if inDegree[i] != 0 {
			return ""
		}
	}
	return result.String()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
