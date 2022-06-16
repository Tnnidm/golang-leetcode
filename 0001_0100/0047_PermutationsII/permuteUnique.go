package permuteUnique

func permuteUnique(nums []int) [][]int {
	var numsCount [21]int
	for _, num := range nums {
		numsCount[num+10]++
	}
	numsLen := len(nums)
	combination := make([]int, numsLen)
	result := [][]int{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == numsLen {
			result = append(result, append([]int{}, combination...))
			return
		}
		for i := -10; i <= 10; i++ {
			if numsCount[i+10] > 0 {
				numsCount[i+10]--
				combination[index] = i
				dfs(index + 1)
				numsCount[i+10]++
			}
		}
	}
	dfs(0)
	return result
}
