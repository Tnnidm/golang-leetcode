package leastInterval

// // 模拟
// func leastInterval(tasks []byte, n int) int {
// 	var taskList [26]struct {
// 		count       int
// 		nextRunTime int
// 	}
// 	for i := range tasks {
// 		taskList[tasks[i]-'A'].count++
// 	}
// 	remain := true
// 	for time := 0; ; time++ {
// 		choice := -1
// 		remain = false
// 		for i := range taskList {
// 			if taskList[i].count > 0 {
// 				remain = true
// 				if taskList[i].nextRunTime <= time && (choice == -1 || taskList[i].count > taskList[choice].count) {
// 					choice = i
// 				}
// 			}
// 		}
// 		if !remain {
// 			return time
// 		} else if choice != -1 {
// 			taskList[choice].nextRunTime = time + n + 1
// 			taskList[choice].count--
// 		}
// 	}
// }

// 构造
func leastInterval(tasks []byte, n int) int {
	maxExecTimes, maxExecCount := 0, 0
	var count [26]int
	for i := range tasks {
		count[tasks[i]-'A']++
		if count[tasks[i]-'A'] > maxExecTimes {
			maxExecTimes = count[tasks[i]-'A']
			maxExecCount = 1
		} else if count[tasks[i]-'A'] == maxExecTimes {
			maxExecCount++
		}
	}
	return max(len(tasks), (maxExecTimes-1)*(n+1)+maxExecCount)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
