package twoSum

// Solution 1: brute-force search
func twoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if (i != j) && (num1+num2 == target) {
				return []int{i, j}
			}
		}
	}
	return nil
}
