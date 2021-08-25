package problem

func rotate(matrix [][]int) {
	quadrats := len(matrix) / 2
	for i := 0; i < quadrats; i++ {
		beg := 0 + i
		end := len(matrix) - 1 - i
		swaps := len(matrix) - 1 - i*2

		for j := 0; j < swaps; j++ {
			topLeft := matrix[beg][beg+j]
			matrix[beg][beg+j] = matrix[end-j][beg]
			matrix[end-j][beg] = matrix[end][end-j]
			matrix[end][end-j] = matrix[beg+j][end]
			matrix[beg+j][end] = topLeft
		}
	}
}
