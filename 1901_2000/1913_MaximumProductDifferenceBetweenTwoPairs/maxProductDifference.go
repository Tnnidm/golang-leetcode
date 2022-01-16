package maxProductDifference

func maxProductDifference(nums []int) int {
	largestNum := 0
	secondLargestNum := 0
	smallestNum := 10000
	secondSmallestNum := 10000
	for _, num := range nums {
		if num > secondLargestNum {
			secondLargestNum = num
			if secondLargestNum > largestNum {
				secondLargestNum, largestNum = largestNum, secondLargestNum
			}
		}
		if num < secondSmallestNum {
			secondSmallestNum = num
			if secondSmallestNum < smallestNum {
				secondSmallestNum, smallestNum = smallestNum, secondSmallestNum
			}
		}
	}
	return largestNum*secondLargestNum - smallestNum*secondSmallestNum
}
