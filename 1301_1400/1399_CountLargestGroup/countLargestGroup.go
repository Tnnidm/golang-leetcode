package countLargestGroup

func countLargestGroup(n int) int {
	bucket := make([]int, 36)
	var count int
	for i := 1; i <= n; i++ {
		count = 0
		for temp := i; temp > 0; temp = temp / 10 {
			count = count + temp%10
		}
		bucket[count-1]++
	}
	maxCount := 0
	maxBucketNum := 0
	for i := 0; i < 36; i++ {
		if bucket[i] > maxBucketNum {
			maxBucketNum = bucket[i]
			maxCount = 1
		} else if bucket[i] == maxBucketNum {
			maxCount++
		}
	}
	return maxCount
}
