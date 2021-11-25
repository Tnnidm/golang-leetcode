package main

func dominantIndex(nums []int) int {
	var first, second, index int
	for i, v := range nums {
		if v > first {
			second = first
			first = v
			index = i
		} else if v > second {
			second = v
		}
	}
	if first >= 2*second {
		return index
	} else {
		return -1
	}
}
