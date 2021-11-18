package findShortestSubArray

// 执行用时：36 ms, 在所有 Go 提交中击败了15.70%的用户
// 内存消耗：6.8 MB, 在所有 Go 提交中击败了61.98%的用户
// 性能太差需要优化
func findShortestSubArray(nums []int) int {
	hashmap := make(map[int][]int)
	maxDegreeNum := nums[0]
	for i := 0; i < len(nums); i++ {
		if status, ok := hashmap[nums[i]]; ok {
			status[0]++
			status[2] = i - status[1] + 1
			if status[0] > hashmap[maxDegreeNum][0] {
				maxDegreeNum = nums[i]
			} else if status[0] == hashmap[maxDegreeNum][0] && status[2] < hashmap[maxDegreeNum][2] {
				maxDegreeNum = nums[i]
			}
		} else {
			hashmap[nums[i]] = []int{1, i, 1}
		}
	}
	return hashmap[maxDegreeNum][2]
}
