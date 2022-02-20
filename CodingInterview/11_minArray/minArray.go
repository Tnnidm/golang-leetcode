package minArray

func minArray(numbers []int) int {
	begin, end := 0, len(numbers)-1
	mid := begin + (end-begin)/2
	for mid != begin {
		if numbers[mid] == numbers[end] && numbers[mid] == numbers[begin] {
			minValue := numbers[mid]
			for i := begin; i < mid; i++ {
				if numbers[i] < minValue {
					minValue = numbers[i]
				}
			}
			if minValue == numbers[mid] {
				begin = mid
				mid = begin + (end-begin)/2
				continue
			} else {
				return minValue
			}
		}
		if numbers[mid] <= numbers[end] {
			end = mid
			mid = begin + (end-begin)/2
			continue
		}
		if numbers[mid] >= numbers[begin] {
			begin = mid
			mid = begin + (end-begin)/2
			continue
		}

	}
	if numbers[begin] >= numbers[end] {
		return numbers[end]
	} else {
		return numbers[begin]
	}
}
