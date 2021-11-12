package findRestaurant

func findRestaurant(list1 []string, list2 []string) []string {
	list1Len := len(list1)
	list2Len := len(list2)
	minIndexSum := 2000
	result := []string{}
	if list1Len > list2Len {
		shortIndex := make(map[string]int, list2Len)
		for i := 0; i < list2Len; i++ {
			shortIndex[list2[i]] = i
		}
		for i := 0; i < list1Len; i++ {
			if index, ok := shortIndex[list1[i]]; ok {
				if i+index < minIndexSum {
					result = []string{list1[i]}
					minIndexSum = i + index
				} else if i+index == minIndexSum {
					result = append(result, list1[i])
				}
			}

		}
	} else {
		shortIndex := make(map[string]int, list1Len)
		for i := 0; i < list1Len; i++ {
			shortIndex[list1[i]] = i
		}
		for i := 0; i < list2Len; i++ {
			if index, ok := shortIndex[list2[i]]; ok {
				if i+index < minIndexSum {
					result = []string{list2[i]}
					minIndexSum = i + index
				} else if i+index == minIndexSum {
					result = append(result, list2[i])
				}
			}
		}
	}
	return result
}
