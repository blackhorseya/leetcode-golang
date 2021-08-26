package problem

import "math"

func setZeroes(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] != 0 {
				continue
			}

			for row2 := 0; row2 < len(matrix); row2++ {
				if matrix[row2][col] != 0 {
					matrix[row2][col] = math.MaxInt64
				}
			}

			for col2 := 0; col2 < len(matrix[0]); col2++ {
				if matrix[row][col2] != 0 {
					matrix[row][col2] = math.MaxInt64
				}
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == math.MaxInt64 {
				matrix[row][col] = 0
			}
		}
	}
}
