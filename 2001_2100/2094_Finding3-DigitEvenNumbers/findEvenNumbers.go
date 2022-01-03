package findEvenNumbers

import "sort"

func findEvenNumbers(digits []int) []int {
	numsMap := map[int]struct{}{}
	numLen := len(digits)
	for i := 0; i < numLen; i++ {
		if digits[i] != 0 {
			for j := 0; j < numLen; j++ {
				if j != i {
					for k := 0; k < numLen; k++ {
						if k != i && k != j && digits[k]%2 == 0 {
							numsMap[digits[i]*100+digits[j]*10+digits[k]] = struct{}{}
						}
					}
				}
			}
		}
	}
	result := make([]int, len(numsMap))
	index := 0
	for k := range numsMap {
		result[index] = k
		index++
	}
	sort.Ints(result)
	return result
}
