package flatten

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	begin, _ := flattenCore(root)
	return begin
}

func flattenCore(root *Node) (*Node, *Node) {
	for temp := root; temp != nil; {
		next := temp.Next
		if temp.Child != nil {
			begin, end := flattenCore(temp.Child)
			temp.Child = nil
			temp.Next = begin
			begin.Prev = temp
			end.Next = next
			if next == nil {
				return root, end
			}
			next.Prev = end
		}
		if next == nil {
			return root, temp
		}
		temp = next
	}
	return root, root
}
