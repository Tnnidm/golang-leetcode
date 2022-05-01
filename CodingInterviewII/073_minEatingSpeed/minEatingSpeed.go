package minEatingSpeed

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	for left <= right {
		mid := left + (right-left)>>1
		if eatingTime(piles, mid) <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func eatingTime(piles []int, speed int) int {
	sum := 0
	for _, pile := range piles {
		sum += (pile-1)/speed + 1
	}
	return sum
}
