package flipAndInvertImage

func flipAndInvertImage(image [][]int) [][]int {
	rLen := len(image)
	cLen := len(image[0])
	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen/2; j++ {
			image[i][j], image[i][cLen-1-j] = image[i][cLen-1-j], image[i][j]
		}
		for j := 0; j < cLen; j++ {
			image[i][j] = 1 - image[i][j]
		}
	}
	return image
}
