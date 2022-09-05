package maximumRobots

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	preSum := make([]int, n+1)
	for i := range runningCosts {
		preSum[i+1] = preSum[i] + runningCosts[i]
	}
	maxLen := 0
	left, right := 0, 0
	maxValueDeQueue := []int{chargeTimes[0]}
	for left < n {
		if maxValueDeQueue[0]+(right-left+1)*(preSum[right+1]-preSum[left]) <= int(budget) {
			maxLen = max(maxLen, right-left+1)
			if right < n-1 {
				right++
				for len(maxValueDeQueue) > 0 && chargeTimes[right] > maxValueDeQueue[len(maxValueDeQueue)-1] {
					maxValueDeQueue = maxValueDeQueue[:len(maxValueDeQueue)-1]
				}
				maxValueDeQueue = append(maxValueDeQueue, chargeTimes[right])
			} else {
				break
			}
		} else {
			if chargeTimes[left] == maxValueDeQueue[0] {
				maxValueDeQueue = maxValueDeQueue[1:]
			}
			left++
			if left < n && left > right {
				right++
				maxValueDeQueue = []int{chargeTimes[right]}
			}
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
