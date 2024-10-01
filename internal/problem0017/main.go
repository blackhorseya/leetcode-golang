package problem0017

// letterCombinations 全排列問題，回溯法
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 建立數字與字母的對應關係
	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	// 回傳結果
	res := []string{}
	// 回溯法
	var backtrack func(int, string)
	backtrack = func(index int, path string) {
		// 終止條件
		if index == len(digits) {
			res = append(res, path)
			return
		}
		// 取得當前數字
		digit := digits[index]
		// 取得當前數字對應的字母
		letters := mapping[digit]
		// 遍歷當前數字對應的字母
		for _, letter := range letters {
			// 遞迴
			backtrack(index+1, path+string(letter))
		}
	}
	backtrack(0, "")
	return res
}
