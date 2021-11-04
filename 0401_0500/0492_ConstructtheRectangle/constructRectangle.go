package constructRectangle

import "math"

// func constructRectangle(area int) []int {
// 	for w := int(math.Sqrt(float64(area))); w >= 1; w-- {
// 		if area/w == 0 {
// 			return []int{area / w, w}
// 		}
// 	}
// 	return []int{area, 1}
// }

// 更省内存的写法
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}
