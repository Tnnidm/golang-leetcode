package nextGreaterElement

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	biggerMap := make(map[int]int, len(nums2))
	var flag bool
	for i := 0; i < len(nums2); i++ {
		flag = true
		for j := i; j < len(nums2); j++ {
			if nums2[i] < nums2[j] {
				biggerMap[nums2[i]] = nums2[j]
				flag = false
				break
			}
			if flag {
				biggerMap[nums2[i]] = -1
			}
		}
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = biggerMap[nums1[i]]
	}
	return nums1
}
