package minArray

func minArray(numbers []int) int {
	numsLen := len(numbers)
	left, right := 0, numsLen-1
	for left < right {
		mid := left + (right-left)>>1
		if mid > 0 && numbers[mid-1] > numbers[mid] {
			return numbers[mid]
		}
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			minValue := numbers[mid]
			for i := mid + 1; i < right; i++ {
				if numbers[i] < minValue {
					minValue = numbers[i]
				}
			}
			if minValue < numbers[mid] {
				return minValue
			} else {
				right = mid
			}
		}
	}
	return numbers[left]
}
