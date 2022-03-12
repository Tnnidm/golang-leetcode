package majorityElement

func majorityElement(nums []int) int {
	var candidate int
	var count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}
