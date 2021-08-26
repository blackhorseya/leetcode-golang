package problem

func dfs(board [][]byte, i, j, pos int, word string) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] == '*' {
		return false
	}

	if word[pos] != board[i][j] {
		return false
	}

	if pos == len(word)-1 {
		return true
	}

	pos++
	temp := board[i][j]
	board[i][j] = '*'

	ret := dfs(board, i-1, j, pos, word) ||
		dfs(board, i, j+1, pos, word) ||
		dfs(board, i+1, j, pos, word) ||
		dfs(board, i, j-1, pos, word)

	board[i][j] = temp

	return ret
}

func exist(board [][]byte, word string) bool {
	pos := 0

	for i, row := range board {
		for j := range row {
			if board[i][j] == word[pos] {
				if ret := dfs(board, i, j, pos, word); ret {
					return true
				}
			}
		}
	}

	return false
}
