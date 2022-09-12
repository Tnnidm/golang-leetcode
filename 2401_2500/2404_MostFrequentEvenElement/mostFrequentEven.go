package mostFrequentEven

func mostFrequentEven(nums []int) int {
	element := -1
	count := map[int]int{-1: 0}
	for _, num := range nums {
		if num&1 == 0 {
			count[num]++
			if count[num] > count[element] || (count[num] == count[element] && num < element) {
				element = num
			}
		}
	}
	return element
}
