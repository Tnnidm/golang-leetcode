package specialArray

func specialArray(nums []int) int {
	var num int
	for x := len(nums); x >= 1; x-- {
		num = 0
		for _, v := range nums {
			if v >= x {
				num++
			}
		}
		if num == x {
			return x
		}
	}
	return -1
}
