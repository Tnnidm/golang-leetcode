package intersect

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	len1, len2 := len(nums1), len(nums2)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len1 && j < len2; {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
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
