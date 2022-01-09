package getMinDistance

func getMinDistance(nums []int, target int, start int) int {
	numsLen := len(nums)
	for dis := 0; ; dis++ {
		if (start+dis < numsLen && nums[start+dis] == target) || (start-dis >= 0 && nums[start-dis] == target) {
			return dis
		}
	}
}
