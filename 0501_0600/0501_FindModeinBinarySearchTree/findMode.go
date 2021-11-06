package findMode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func findMode(root *TreeNode) []int {
// 	var count, maxCount int
// 	base := math.MinInt
// 	result := []int{}
// 	midOrderSearch(root, &base, &count, &maxCount, &result)
// 	return result
// }

// func midOrderSearch(node *TreeNode, base *int, count *int, maxCount *int, result *[]int) {
// 	if node != nil {
// 		midOrderSearch(node.Left, base, count, maxCount, result)
// 		update(node.Val, base, count, maxCount, result)
// 		midOrderSearch(node.Right, base, count, maxCount, result)
// 	}
// }

// func update(x int, base *int, count *int, maxCount *int, result *[]int) {
// 	if x > *base {
// 		*base = x
// 		*count = 1
// 	} else if x == *base {
// 		*count += 1
// 	}

// 	if *count > *maxCount {
// 		*maxCount = *count
// 		*result = []int{x}
// 	} else if *count == *maxCount {
// 		*result = append(*result, *base)
// 	}
// }

// 更简单也消耗更小的写法
func findMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
