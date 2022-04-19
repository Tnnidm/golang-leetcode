package findTarget

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	arr := []int{}
	dfs(root, &arr)
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] == k {
			return true
		} else if arr[left]+arr[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}

func dfs(root *TreeNode, arr *[]int) {
	if root != nil {
		dfs(root.Left, arr)
		*arr = append(*arr, root.Val)
		dfs(root.Right, arr)
	}
}
