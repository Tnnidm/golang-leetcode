package matrixReshape

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rLen := len(mat)
	cLen := len(mat[0])
	if rLen*cLen != r*c {
		return mat
	} else {
		var arr = make([][]int, r)
		// 然后分别对其中的进行make
		for i := range arr {
			arr[i] = make([]int, c)
		}
		for i := 0; i < rLen*cLen; i++ {
			arr[i/c][i%c] = mat[i/cLen][i%cLen]
		}
		return arr
	}
}
