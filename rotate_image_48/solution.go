package rotate_image_48

func rotate(matrix [][]int) {
	// transpose the matrix in-place, by turning it's rows into columns
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row of the transposed matrix to get the 90deg rotated matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-1-j] =
				matrix[i][len(matrix[i])-1-j], matrix[i][j]
		}
	}
}
