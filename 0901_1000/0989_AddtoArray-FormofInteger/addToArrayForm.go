package addToArrayForm

func addToArrayForm(num []int, k int) []int {
	addOne := 0
	i := len(num) - 1
	for k != 0 || addOne == 1 {
		if i >= 0 {
			num[i] = num[i] + k%10 + addOne
			if num[i] >= 10 {
				num[i] = num[i] - 10
				addOne = 1
			} else {
				addOne = 0
			}
			i--
		} else {
			if k%10+addOne >= 10 {
				num = append([]int{k%10 + addOne - 10}, num...)
				addOne = 1
			} else {
				num = append([]int{k%10 + addOne}, num...)
				addOne = 0
			}
		}
		k = k / 10
	}
	return num
}
