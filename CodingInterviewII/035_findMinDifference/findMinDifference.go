package findMinDifference

import "strconv"

func findMinDifference(timePoints []string) int {
	var time [1440]bool
	for _, timePoint := range timePoints {
		if time[transferToMinuteIndex(timePoint)] {
			return 0
		}
		time[transferToMinuteIndex(timePoint)] = true
	}
	lastTime, minDifference := -1440, 1440
	for i := 0; i < 1440; i++ {
		if time[i] {
			minDifference = min(minDifference, i-lastTime)
			lastTime = i
		}
	}
	gapDayDifference := 1
	for i := 0; !time[i]; i++ {
		gapDayDifference++
	}
	for i := 1439; !time[i]; i-- {
		gapDayDifference++
	}
	return min(minDifference, gapDayDifference)
}

func transferToMinuteIndex(time string) int {
	hour, _ := strconv.Atoi(time[:2])
	minute, _ := strconv.Atoi(time[3:])
	return 60*hour + minute
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
