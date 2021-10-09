package twoSum

func twoSum(numbers []int, target int) []int {
	smaller, bigger := 0, len(numbers)-1
	var sum int
	for smaller < bigger {
		sum = numbers[smaller] + numbers[bigger]
		if sum > target {
			bigger -= 1
		} else if sum < target {
			smaller += 1
		} else {
			return []int{smaller + 1, bigger + 1}
		}
	}
	return []int{}
}
