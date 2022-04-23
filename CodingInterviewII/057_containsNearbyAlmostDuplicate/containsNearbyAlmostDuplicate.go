package containsNearbyAlmostDuplicate

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	numsLen := len(nums)
	bucket := make(map[int]int, k+1)
	if t == 0 {
		for i := 0; i < numsLen; i++ {
			if i >= k+1 {
				delete(bucket, nums[i-k-1])
			}
			if _, ok := bucket[nums[i]]; ok {
				return true
			}
			bucket[nums[i]] = nums[i]
		}
	} else {
		for i := 0; i < numsLen; i++ {
			if i >= k+1 {
				delete(bucket, getBucketIndex(nums[i-k-1], t))
			}
			bucketIndex := getBucketIndex(nums[i], t)
			if _, ok := bucket[bucketIndex]; ok {
				return true
			} else {
				if _, ok := bucket[bucketIndex-1]; ok && abs(bucket[bucketIndex-1]-nums[i]) <= t {
					return true
				}
				if _, ok := bucket[bucketIndex+1]; ok && abs(bucket[bucketIndex+1]-nums[i]) <= t {
					return true
				}
				bucket[bucketIndex] = nums[i]
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getBucketIndex(x int, t int) int {
	if x < 0 {
		return x/t - 1
	}
	return x / t
}
