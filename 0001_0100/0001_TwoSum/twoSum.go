package twoSum

// Solution 1: brute-force search
// func twoSum(nums []int, target int) []int {
// 	for i, num1 := range nums {
// 		for j, num2 := range nums {
// 			if (i != j) && (num1+num2 == target) {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

// Solution 2: map
func twoSum(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i, num := range nums {
		if j, ok := m1[target-num]; ok {
			return []int{j, i}
		}
		m1[num] = i
	}
	return nil
}
