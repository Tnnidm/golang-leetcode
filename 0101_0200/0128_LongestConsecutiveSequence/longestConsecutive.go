package longestConsecutive

func longestConsecutive(nums []int) int {
	numsLen := len(nums)
	numsExist := make(map[int]bool, numsLen)
	for _, num := range nums {
		numsExist[num] = true
	}
	result := 0
	for num := range nums {
		if numsExist[num] {
			numsExist[num] = false
			length := 1
			for i := num - 1; numsExist[i]; i-- {
				length++
				numsExist[i] = false
			}
			for i := num + 1; numsExist[i]; i++ {
				length++
				numsExist[i] = false
			}
			if length > result {
				result = length
			}
		}
	}
	return result
}
