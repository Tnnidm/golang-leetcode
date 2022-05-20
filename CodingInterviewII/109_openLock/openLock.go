package openLock

// 双向BFS
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	dead := make(map[string]bool)
	for _, deadend := range deadends {
		dead[deadend] = true
	}
	if dead["0000"] {
		return -1
	}
	visited1 := map[string]bool{"0000": true}
	visited2 := map[string]bool{target: true}
	queue1 := []string{"0000"}
	queue2 := []string{target}
	queueLen1 := len(queue1)
	queueLen2 := len(queue2)
	step := 0
	for queueLen1 != 0 && queueLen2 != 0 {
		if queueLen1 > queueLen2 {
			queue1, queue2 = queue2, queue1
			queueLen1, queueLen2 = queueLen2, queueLen1
			visited1, visited2 = visited2, visited1
		}
		queue3 := []string{}
		for _, node := range queue1 {
			for _, neighbor := range getNeighbor(node) {
				if visited2[neighbor] {
					return step + 1
				}
				if !visited1[neighbor] && !dead[neighbor] {
					visited1[neighbor] = true
					queue3 = append(queue3, neighbor)
				}
			}
		}
		queue1 = queue3
		queueLen1 = len(queue1)
		step++
	}
	return -1
}

func getNeighbor(word string) []string {
	neighbors := make([]string, 8)
	wordByte := []byte(word)
	for i := 0; i < 4; i++ {
		temp := wordByte[i]
		if wordByte[i] == '0' {
			wordByte[i] = '9'
			neighbors[i<<1] = string(wordByte)
			wordByte[i] = '1'
			neighbors[i<<1+1] = string(wordByte)
		} else if wordByte[i] == '9' {
			wordByte[i] = '0'
			neighbors[i<<1] = string(wordByte)
			wordByte[i] = '8'
			neighbors[i<<1+1] = string(wordByte)
		} else {
			wordByte[i] = temp - 1
			neighbors[i<<1] = string(wordByte)
			wordByte[i] = temp + 1
			neighbors[i<<1+1] = string(wordByte)
		}
		wordByte[i] = temp
	}
	return neighbors
}
