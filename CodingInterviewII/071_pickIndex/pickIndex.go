package pickIndex

import "math/rand"

type Solution struct {
	sum    []int
	length int
}

func Constructor(w []int) Solution {
	wLen := len(w)
	for i := 1; i < wLen; i++ {
		w[i] += w[i-1]
	}
	return Solution{w, wLen}
}

func (this *Solution) PickIndex() int {
	randomNum := rand.Intn(this.sum[this.length-1]) + 1
	left, right := 0, this.length-1
	for left <= right {
		mid := left + (right-left)>>1
		if this.sum[mid] >= randomNum {
			if mid == 0 || this.sum[mid-1] < randomNum {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return this.length
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
