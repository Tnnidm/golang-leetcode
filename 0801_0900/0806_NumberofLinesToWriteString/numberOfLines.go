package numberOfLines

func numberOfLines(widths []int, s string) []int {
	remainUnit := 100
	rowNum := 1
	var width int
	for _, char := range s {
		width = widths[char-'a']
		if width <= remainUnit {
			remainUnit -= width
		} else {
			remainUnit = 100 - width
			rowNum++
		}
	}
	return []int{rowNum, 100 - remainUnit}
}
