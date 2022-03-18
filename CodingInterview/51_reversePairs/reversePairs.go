package reversePairs

func reversePairs(nums []int) int {
	count := 0
	reverse(&nums, 0, len(nums)-1, &count)
	return count
}

func reverse(nums *[]int, begin, end int, count *int) {
	if begin == end {
		return
	}
	mid := begin + (end-begin)>>1
	if begin < mid {
		reverse(nums, begin, mid, count)
	}
	if mid+1 < end {
		reverse(nums, mid+1, end, count)
	}
	assistArray := make([]int, end-begin+1)
	p1, p2 := mid, end
	index := end - begin
	for p1 >= begin && p2 >= mid+1 {
		if (*nums)[p1] > (*nums)[p2] {
			assistArray[index] = (*nums)[p1]
			p1--
			*count += p2 - mid
		} else {
			assistArray[index] = (*nums)[p2]
			p2--
		}
		index--
	}
	for p1 >= begin {
		assistArray[index] = (*nums)[p1]
		index--
		p1--
	}
	for p2 >= mid+1 {
		assistArray[index] = (*nums)[p2]
		index--
		p2--
	}
	for i := begin; i <= end; i++ {
		(*nums)[i] = assistArray[i-begin]
	}
}
