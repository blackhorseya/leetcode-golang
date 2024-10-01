package problem0022

// generateParenthesis 全排列問題，使用回溯法
func generateParenthesis(n int) []string {
	var res []string
	backtrack("", n, n, &res)
	return res
}

// backtrack 遞迴生成合法的括號組合
// left 代表還能放多少個 '('，right 代表還能放多少個 ')'
func backtrack(state string, left, right int, res *[]string) {
	// 如果左右括號都用完了，則當前 state 為一個合法解
	if left == 0 && right == 0 {
		*res = append(*res, state)
		return
	}

	// 如果還有左括號可以放，則放一個左括號
	if left > 0 {
		backtrack(state+"(", left-1, right, res)
	}

	// 如果右括號比左括號多，則放一個右括號
	if right > left {
		backtrack(state+")", left, right-1, res)
	}
}
