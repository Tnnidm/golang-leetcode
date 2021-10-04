package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// A BFS algorithm
	row := []*TreeNode{root}
	for len(row) > 0 {
		nextRow := []*TreeNode{}
		for i := 0; i < len(row); i++ {
			if row[i].Left == nil && row[i].Right == nil {
				if row[i].Val == targetSum {
					return true
				}
			} else {
				if row[i].Left != nil {
					row[i].Left.Val += row[i].Val
					nextRow = append(nextRow, row[i].Left)
				}
				if row[i].Right != nil {
					row[i].Right.Val += row[i].Val
					nextRow = append(nextRow, row[i].Right)
				}
			}
		}
		row = nextRow
	}
	return false
}

// // Other's method
// func hasPathSum(root *TreeNode, targetSum int) bool {
//     if root ==nil{
//         return false
//     }
//     if root.Left==nil && root.Right ==nil{
//         return root.Val == targetSum
//     }
//     if root.Left ==nil{
//         return hasPathSum(root.Right,targetSum-root.Val)
//     }
//     if root.Right ==nil{
//         return hasPathSum(root.Left,targetSum-root.Val)
//     }

//     return hasPathSum(root.Left,targetSum-root.Val)|| hasPathSum(root.Right,targetSum-root.Val)
// }
