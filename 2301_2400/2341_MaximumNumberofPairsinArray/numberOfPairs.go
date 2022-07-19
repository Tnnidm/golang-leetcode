package numberOfPairs

func numberOfPairs(nums []int) []int {
	result := []int{0, 0}
	exits := map[int]bool{}
	for _, num := range nums {
		if exits[num] {
			exits[num] = false
			result[1]--
			result[0]++
		} else {
			exits[num] = true
			result[1]++
		}
	}
	return result
}
