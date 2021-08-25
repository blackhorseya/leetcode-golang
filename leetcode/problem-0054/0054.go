package problem

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	ret := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for {
		// right
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// down
		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// left
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// up
		for i := bottom; i >= top; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return ret
}
