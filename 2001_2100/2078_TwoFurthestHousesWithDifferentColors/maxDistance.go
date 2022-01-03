package maxDistance

func maxDistance(colors []int) int {
	begin := 0
	end := len(colors) - 1
	if colors[begin] != colors[end] {
		return end
	} else {
		for begin == 0 || colors[begin] == colors[begin-1] {
			begin++
		}
		for end == len(colors)-1 || colors[end] == colors[end+1] {
			end--
		}
		return len(colors) - 1 - min(begin, len(colors)-1-end)
	}
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
