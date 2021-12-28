package shuffle

func shuffle(nums []int, n int) []int {
	var temp int
	for i := n; i < 2*n; i++ {
		temp = nums[i]
		for j := i; j > 2*(i-n)+1; j-- {
			nums[j] = nums[j-1]
		}
		nums[2*(i-n)+1] = temp
	}
	return nums
}
