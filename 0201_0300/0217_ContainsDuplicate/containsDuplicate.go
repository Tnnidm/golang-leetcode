package containsDuplicate

// 执行用时：16 ms, 在所有 Go 提交中击败了95.26%的用户
// 内存消耗：7.3 MB, 在所有 Go 提交中击败了72.97%的用户
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for i := range nums {
		_, found := m[nums[i]]
		if found {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}
