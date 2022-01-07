package maximumUnits

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] >= boxTypes[j][1]
	})
	// 单元数
	unit := 0
	for _, v := range boxTypes {
		if truckSize > v[0] {
			unit += v[1] * v[0]
			truckSize = truckSize - v[0]
		} else {
			unit += v[1] * truckSize
			truckSize = 0
			break
		}
	}
	return unit
}
