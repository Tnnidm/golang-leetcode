package findMedianSortedArrays

// 二分查找法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	bias1, bias2 := 0, 0
	for {
		if bias1 == len(nums1) {
			return nums2[bias2+k-1]
		}
		if bias2 == len(nums2) {
			return nums1[bias1+k-1]
		}
		if k == 1 {
			return min(nums1[bias1], nums2[bias2])
		}
		half := k / 2
		index1 := min(bias1+half, len(nums1)) - 1
		index2 := min(bias2+half, len(nums2)) - 1
		if nums1[index1] <= nums2[index2] {
			k -= (index1 - bias1 + 1)
			bias1 = index1 + 1
		} else {
			k -= (index2 - bias2 + 1)
			bias2 = index2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
