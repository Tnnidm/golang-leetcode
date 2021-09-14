package merge

// // my mthod
// func merge(nums1 []int, m int, nums2 []int, n int) []int {
// 	if n != 0 {
// 		if m == 0 {
// 			copy(nums1, nums2)
// 		} else {
// 			nIndex := 0
// 			for mIndex := 0; mIndex < m+n-1; mIndex++ {
// 				if nums2[nIndex] <= nums1[mIndex] {
// 					// nums2[nIndex] should be place in mIndex
// 					copy(nums1[mIndex+1:m+n], nums1[mIndex:m+n-1])
// 					nums1[mIndex] = nums2[nIndex]
// 					nIndex += 1
// 				}
// 				if nIndex == n {
// 					break
// 				}
// 				if mIndex == m-1+nIndex {
// 					copy(nums1[mIndex+1:], nums2[nIndex:])
// 					break
// 				}

// 			}
// 		}
// 	}
// }

// other's method: 倒序法，从后往前选最大的
// more neat
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p-1] = nums1[m-1]
			m--
		} else {
			nums1[p-1] = nums2[n-1]
			n--
		}
	}
	// 可能nums2还有剩余
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
