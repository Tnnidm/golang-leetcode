package findKthPositive

func findKthPositive(arr []int, k int) int {
	index := 0
	for i := 1; ; i++ {
		if i != arr[index] {
			k--
			if k == 0 {
				return i
			}
		} else if index < len(arr)-1 {
			index++
		}
	}
}
