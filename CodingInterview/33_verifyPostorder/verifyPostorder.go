package verifyPostorder

func verifyPostorder(postorder []int) bool {
	postLen := len(postorder)
	if postLen < 2 {
		return true
	}
	leftHead := postLen - 2
	for leftHead >= 0 && postorder[leftHead] > postorder[postLen-1] {
		leftHead--
	}
	for i := 0; i <= leftHead; i++ {
		if postorder[i] > postorder[postLen-1] {
			return false
		}
	}
	return verifyPostorder(postorder[:leftHead+1]) && verifyPostorder(postorder[leftHead+1:postLen-1])
}
