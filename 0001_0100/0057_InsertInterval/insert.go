package insert

func insert(intervals [][]int, newInterval []int) [][]int {
	intervalNum := len(intervals)
	if intervalNum == 0 {
		return [][]int{newInterval}
	}
	left, right := 0, intervalNum
	for left < right {
		mid := (left + right) >> 1
		if intervals[mid][0] > newInterval[0] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	right = left
	for left > 0 && intervals[left-1][1] >= newInterval[0] {
		newInterval[0] = min(newInterval[0], intervals[left-1][0])
		newInterval[1] = max(newInterval[1], intervals[left-1][1])
		left--
	}
	for right < intervalNum && intervals[right][0] <= newInterval[1] {
		newInterval[1] = max(newInterval[1], intervals[right][1])
		newInterval[0] = min(newInterval[0], intervals[right][0])
		right++
	}
	return append(intervals[:left], append([][]int{newInterval}, intervals[right:]...)...)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
