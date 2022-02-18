package typicalSort

func mergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	middle := arrLen / 2
	return merge(mergeSort(arr[:middle]), mergeSort(arr[middle:]))
}

func merge(part1 []int, part2 []int) []int {
	part1Len := len(part1)
	part2Len := len(part2)
	result := make([]int, part1Len+part2Len)
	i, j := 0, 0
	for i < part1Len && j < part2Len {
		if part1[i] <= part2[j] {
			result[i+j] = part1[i]
			i++
		} else {
			result[i+j] = part2[j]
			j++
		}
	}
	for i < part1Len {
		result[i+j] = part1[i]
		i++
	}
	for j < part2Len {
		result[i+j] = part2[j]
		j++
	}
	return result
}
