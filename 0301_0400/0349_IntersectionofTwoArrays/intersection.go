package intersection

import "sort"

// func intersection(nums1 []int, nums2 []int) []int {
// 	set := map[int]struct{}{}
// 	for _, num := range nums1 {
// 		if _, ok := set[num]; !ok {
// 			set[num] = struct{}{}
// 		}
// 	}
// 	result := []int{}
// 	resultSet := map[int]struct{}{}
// 	for _, num := range nums2 {
// 		if _, ok := set[num]; ok {
// 			if _, ok := resultSet[num]; !ok {
// 				result = append(result, num)
// 				resultSet[num] = struct{}{}
// 			}
// 		}
// 	}
// 	return result
// }

// 解法2：先转换程有序数组
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	len1, len2 := len(nums1), len(nums2)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len1 && j < len2; {
		if nums1[i] == nums2[j] {
			if len(result) == 0 || nums1[i] != result[len(result)-1] {
				result = append(result, nums1[i])
			}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
