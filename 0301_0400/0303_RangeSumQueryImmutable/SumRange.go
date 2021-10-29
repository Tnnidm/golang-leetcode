package SumRange

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{}
	na.sums = append(na.sums, 0)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		na.sums = append(na.sums, sum)
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
