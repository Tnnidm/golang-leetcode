package twoSum

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	indexMap := make(map[int]int, numsLen)
	for _, num := range nums {
		if anothernum, ok := indexMap[num]; ok {
			return []int{anothernum, num}
		} else {
			indexMap[target-num] = num
		}
	}
	return []int{}
}
