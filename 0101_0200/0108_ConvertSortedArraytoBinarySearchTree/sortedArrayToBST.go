package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	} else {
		mid := numsLen / 2
		t := &TreeNode{Val: nums[mid]}
		t.Left = sortedArrayToBST(nums[:mid])
		t.Right = sortedArrayToBST(nums[mid+1:])
		return t
	}
}
