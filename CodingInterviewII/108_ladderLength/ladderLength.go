package ladderLength

// // 单向BFS
// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	// 建图
// 	wordLen := len(beginWord)
// 	wordList = append(wordList, beginWord)
// 	wordNum := len(wordList)
// 	connection := make([][]int, wordNum)
// 	for i := 0; i < wordNum; i++ {
// 		connection[i] = []int{}
// 	}
// 	for i := 0; i < wordNum; i++ {
// 		for j := i + 1; j < wordNum; j++ {
// 			if linkExist(&wordList[i], &wordList[j], &wordLen) {
// 				connection[i] = append(connection[i], j)
// 				connection[j] = append(connection[j], i)
// 			}
// 		}
// 	}
// 	// BFS搜索
// 	// wordNum-1就是beginword
// 	queue := append([]int{}, connection[wordNum-1]...)
// 	visit := make([]bool, wordNum)
// 	visit[wordNum-1] = true
// 	result := 2
// 	for len(queue) != 0 {
// 		nextQueue := []int{}
// 		for _, node := range queue {
// 			if wordList[node] == endWord {
// 				return result
// 			}
// 			visit[node] = true
// 			for _, neighbor := range connection[node] {
// 				if !visit[neighbor] {
// 					nextQueue = append(nextQueue, neighbor)
// 				}
// 			}
// 		}
// 		queue = nextQueue
// 		result++
// 	}
// 	return 0
// }

// // 双向BFS
// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	// 建图
// 	wordLen := len(beginWord)
// 	wordList = append(wordList, beginWord)
// 	wordNum := len(wordList)
// 	endIndex := -1
// 	connection := make([][]int, wordNum)
// 	for i := 0; i < wordNum; i++ {
// 		connection[i] = make([]int, 0, 8)
// 	}
// 	for i := 0; i < wordNum; i++ {
// 		if wordList[i] == endWord {
// 			endIndex = i
// 		}
// 		for j := i + 1; j < wordNum; j++ {
// 			if linkExist(&wordList[i], &wordList[j], &wordLen) {
// 				connection[i] = append(connection[i], j)
// 				connection[j] = append(connection[j], i)
// 			}
// 		}
// 	}
// 	if endIndex == -1 {
// 		return 0
// 	}
// 	// BFS搜索
// 	// wordNum-1就是beginword
// 	startQueue := []int{wordNum - 1}
// 	endQueue := []int{endIndex}
// 	startQueueLen := len(startQueue)
// 	endQueueLen := len(endQueue)
// 	visit := make([]int, wordNum)
// 	result := 0
// 	for startQueueLen != 0 && endQueueLen != 0 {
// 		for i := 0; i < startQueueLen; i++ {
// 			node := startQueue[0]
// 			startQueue = startQueue[1:]
// 			visit[node] = 1
// 			for _, neighbor := range connection[node] {
// 				if visit[neighbor] == 2 {
// 					return result + 1
// 				} else if visit[neighbor] == 0 {
// 					startQueue = append(startQueue, neighbor)
// 				}
// 			}
// 		}
// 		result++
// 		startQueueLen = len(startQueue)
// 		for i := 0; i < endQueueLen; i++ {
// 			node := endQueue[0]
// 			endQueue = endQueue[1:]
// 			visit[node] = 2
// 			for _, neighbor := range connection[node] {
// 				if visit[neighbor] == 1 {
// 					return result + 1
// 				} else if visit[neighbor] == 0 {
// 					endQueue = append(endQueue, neighbor)
// 				}
// 			}
// 		}
// 		result++
// 		endQueueLen = len(endQueue)
// 	}
// 	return 0
// }

// func linkExist(word1 *string, word2 *string, wordLen *int) bool {
// 	count := 0
// 	for i := 0; i < *wordLen; i++ {
// 		if (*word1)[i] != (*word2)[i] {
// 			count++
// 		}
// 		if count > 1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// 上面建图的办法的复杂度是O(n^2*m),没有利用之差一位这个特点，可以使用虚拟结点法把时间复杂度减到O（nm）
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordLen := len(beginWord)
	wordList = append(wordList, beginWord)
	wordNum := len(wordList)
	virtualWordMap := map[string]([]int){}
	for i := 0; i < wordNum; i++ {
		word := []byte(wordList[i])
		for j := 0; j < wordLen; j++ {
			temp := word[j]
			word[j] = '*'
			virtualWordMap[string(word)] = append(virtualWordMap[string(word)], i)
			word[j] = temp
		}
	}
	connection := make([][]int, wordNum)
	endIndex := -1
	for i := 0; i < wordNum; i++ {
		connection[i] = []int{}
		if wordList[i] == endWord {
			endIndex = i
		}
		word := []byte(wordList[i])
		for j := 0; j < wordLen; j++ {
			temp := word[j]
			word[j] = '*'
			for _, index := range virtualWordMap[string(word)] {
				if i != index {
					connection[i] = append(connection[i], index)
				}
			}
			word[j] = temp
		}
	}
	if endIndex == -1 {
		return 0
	}
	// BFS搜索
	// wordNum-1就是beginword
	startQueue := []int{wordNum - 1}
	endQueue := []int{endIndex}
	startQueueLen := len(startQueue)
	endQueueLen := len(endQueue)
	visit := make([]int, wordNum)
	result := 0
	for startQueueLen != 0 && endQueueLen != 0 {
		for i := 0; i < startQueueLen; i++ {
			node := startQueue[0]
			startQueue = startQueue[1:]
			visit[node] = 1
			for _, neighbor := range connection[node] {
				if visit[neighbor] == 2 {
					return result + 1
				} else if visit[neighbor] == 0 {
					startQueue = append(startQueue, neighbor)
				}
			}
		}
		result++
		startQueueLen = len(startQueue)
		for i := 0; i < endQueueLen; i++ {
			node := endQueue[0]
			endQueue = endQueue[1:]
			visit[node] = 2
			for _, neighbor := range connection[node] {
				if visit[neighbor] == 1 {
					return result + 1
				} else if visit[neighbor] == 0 {
					endQueue = append(endQueue, neighbor)
				}
			}
		}
		result++
		endQueueLen = len(endQueue)
	}
	return 0
}
