package kWeakestRows

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	rowNum := len(mat)
	colNum := len(mat[0])
	matOnes := make([][]int, rowNum)
	for i := 0; i < rowNum; i++ {
		ones := 0
		for j := 0; j < colNum; j++ {
			ones = ones + mat[i][j]
		}
		matOnes[i] = []int{i, ones}
	}
	sort.SliceStable(matOnes, func(i, j int) bool {
		return matOnes[i][1] < matOnes[j][1] || matOnes[i][1] == matOnes[j][1] && i < j
	})
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = matOnes[i][0]
	}
	return result
}
