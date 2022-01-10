package twoOutOfThree

type exist struct {
	inNums1 int
	inNums2 int
	inNums3 int
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	bucket := make([]exist, 101)
	for _, v := range nums1 {
		bucket[v].inNums1 = 1
	}
	for _, v := range nums2 {
		bucket[v].inNums2 = 1
	}
	for _, v := range nums3 {
		bucket[v].inNums3 = 1
	}
	result := []int{}
	for i := 1; i <= 100; i++ {
		if bucket[i].inNums1+bucket[i].inNums2+bucket[i].inNums3 >= 2 {
			result = append(result, i)
		}
	}
	repeatTimes := 0
	for equal(result, nums1) || equal(result, nums2) || equal(result, nums3) {
		if repeatTimes == len(result) {
			break
		}
		result = append(result[1:], result[0])
		repeatTimes++
	}
	return result
}

func equal(x, y []int) bool {
	xLen := len(x)
	yLen := len(y)
	if xLen != yLen {
		return false
	}
	for i := 0; i < xLen; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
