package findTarget

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	arr := []int{}
	midOrderSearch(root, &arr)
	small, large := 0, len(arr)-1
	for small < large {
		if arr[small]+arr[large] < k {
			small++
		} else if arr[small]+arr[large] > k {
			large--
		} else {
			return true
		}
	}
	return false
}

func midOrderSearch(root *TreeNode, p *[]int) {
	if root != nil {
		midOrderSearch(root.Left, p)
		*p = append(*p, root.Val)
		midOrderSearch(root.Right, p)
	}
}
